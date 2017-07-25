
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE request ALTER COLUMN terminate_place TYPE varchar(200);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE request ALTER COLUMN terminate_place TYPE varchar(100);
