-- +goose Up
create table if not exists users (
    id bigserial,
    name varchar not null,
    email varchar not null,
    role varchar not null,
    password_hash varchar,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down
drop table if exists users;