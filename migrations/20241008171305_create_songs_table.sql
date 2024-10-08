-- +goose Up
-- +goose StatementBegin
create table songs(
    id serial primary key,
    group varchar(255) not null,
    song_title varchar(255) not null,
    release_date date not null,
    text text not null,
    link varchar(255) not null,
    created_at timestampz not null default now(),
    updated_at timestampz not null default now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table songs;
-- +goose StatementEnd
