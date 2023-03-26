-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION pg_trgm;
CREATE INDEX users_search_index_gin on users using gin (first_name gin_trgm_ops);
CREATE INDEX users_search_index_second_name_gin on users using gin (second_name gin_trgm_ops);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd
