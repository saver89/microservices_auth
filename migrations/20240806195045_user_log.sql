-- +goose Up
create table if not exists user_logs (
    id bigserial primary key,
    user_id bigint not null,
    log text not null,
    created_at timestamp not null default now()
);

-- +goose Down
drop table if exists user_logs;
