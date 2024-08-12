-- +goose Up
create table if not exists users (
    id bigserial primary key,
    name text not null,
    email text not null,
    role text not null,
    password_hash text,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);

-- +goose Down
drop table if exists users;