-- +goose Up
-- +goose StatementBegin
CREATE TABLE songs(
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_title VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    text TEXT NOT NULL,
    link VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE songs;
-- +goose StatementEnd
