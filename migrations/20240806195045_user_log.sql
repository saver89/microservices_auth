-- +goose Up
create table if not exists user_logs (
    id bigserial,
    user_id bigint not null,
    log varchar not null,
    created_at timestamp not null default now(),
);

-- +goose Down
drop table if exists user_logs;
