-- +goose Up
ALTER TABLE feeds ADD COLUMN last_time_fetched TIMESTAMP;

-- +goose Down
ALTER TABLE feeds DROP COLUMN last_time_fetched;