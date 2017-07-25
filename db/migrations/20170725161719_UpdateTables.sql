
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE request ADD COLUMN currency varchar (256);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE request DROP COLUMN currency;
