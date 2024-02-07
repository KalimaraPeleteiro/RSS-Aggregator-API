-- +goose Up
ALTER TABLE users ADD COLUMN password VARCHAR(64) NOT NULL DEFAULT '1234';

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;