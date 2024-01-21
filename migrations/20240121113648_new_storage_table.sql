-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS storage (
    id BIGSERIAL PRIMARY KEY,
	name VARCHAR NOT NULL,
    is_available BOOLEAN NOT NULL DEFAULT TRUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS storage;
-- +goose StatementEnd
