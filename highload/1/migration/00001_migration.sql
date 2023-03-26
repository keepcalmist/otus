-- +goose Up
-- +goose StatementBegin
create table if not exists users (
    id bigserial  primary key,
    first_name  text not null,
    second_name text not null,
    age        smallint not null,
    biography  text not null,
    city       text not null,
    password   bytea not null
);

alter table users alter column biography set default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
