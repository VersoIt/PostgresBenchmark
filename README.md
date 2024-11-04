# Benchmark для PostgreSQL

Параметры Benchmark настраиваются в файле ```./config.yml```

## Config.yml
```
dsn - строка подключения к БД
sql-query - оцениваемый sql-запрос
test-duration-millis - время проведения измерения RPS в миллисекундах
workers-count - количество одновременно запускаемых горутин
```

## Запуск
```
go run main.go
```
