package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"postgresbenchmark/pkg/cfg"
	"postgresbenchmark/pkg/db"
	"syscall"
	"time"
)

func main() {
	config := cfg.Get()
	database, err := sqlx.Open("postgres", config.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(database)

	ctx, cancel := context.WithTimeout(context.Background(), config.TestDurationMillis*time.Millisecond)
	defer cancel()

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-exit:
			cancel()
		case <-ctx.Done():
			return
		}
	}()

	dbBenchmark := db.NewBenchmark(database, config.WorkersCount)
	requestsCount := dbBenchmark.Exec(ctx, config.SqlQuery, config.WorkersCount)

	select {
	case <-ctx.Done():
		fmt.Printf("RPS: %d\n", requestsCount*1000/int64(config.TestDurationMillis))
		return
	}
}
