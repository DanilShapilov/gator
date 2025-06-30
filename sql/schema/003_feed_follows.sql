-- +goose Up
CREATE TABLE feed_follows (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    feed_id INTEGER NOT NULL REFERENCES feeds ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;