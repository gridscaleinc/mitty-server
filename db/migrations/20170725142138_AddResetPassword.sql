
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS reset_passwords (
    id            serial primary key,
    email         varchar(1024) NOT NULL,
    token         varchar(1024) NOT NULL,
    expire        timestamp NOT NULL,
    created       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE reset_passwords;
