-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS benchmark(
    id SERIAL PRIMARY KEY,
    description VARCHAR(255),
    date DATE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS benchmark;
-- +goose StatementEnd
