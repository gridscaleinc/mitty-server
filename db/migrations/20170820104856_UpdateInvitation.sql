
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE invitation ALTER COLUMN id_of_type TYPE int8 USING(0);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE invitation ALTER COLUMN id_of_type TYPE varchar(255);
