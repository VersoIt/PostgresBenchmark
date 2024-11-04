package cfg

import (
	"github.com/go-yaml/yaml"
	"io"
	"os"
	"sync"
	"time"
)

type Config struct {
	Dsn                string        `yaml:"dsn"`
	SqlQuery           string        `yaml:"sql-query"`
	TestDurationMillis time.Duration `yaml:"test-duration-millis"`
	WorkersCount       int           `yaml:"workers-count"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		f, err := os.Open("config.yml")
		if err != nil {
			panic(err)
		}
		defer func(f *os.File) {
			_ = f.Close()
		}(f)

		bytes, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}

		if err = yaml.Unmarshal(bytes, &cfg); err != nil {
			panic(err)
		}
	})

	return cfg
}
