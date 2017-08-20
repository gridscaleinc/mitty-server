
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE namecard ADD COLUMN business_logo_id int8;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE namecard DROP COLUMN business_logo_id;
