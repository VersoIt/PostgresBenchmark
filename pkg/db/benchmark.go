package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
	"sync/atomic"
)

type Benchmark struct {
	db      *sqlx.DB
	workers int
}

func NewBenchmark(db *sqlx.DB, workers int) *Benchmark {
	return &Benchmark{db: db, workers: workers}
}

func (db *Benchmark) Exec(ctx context.Context, query string, workersCount int) int64 {
	requestsCount := atomic.Int64{}

	wg := sync.WaitGroup{}
	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				_, err := db.db.ExecContext(ctx, query)
				if err != nil {
					log.Printf("Error executing query %s: %s", query, err)
					return
				}

				requestsCount.Add(1)
			}
		}()
	}

	wg.Wait()
	return requestsCount.Load()
}
