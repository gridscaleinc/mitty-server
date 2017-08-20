
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE namecard ADD COLUMN name varchar(200);
ALTER TABLE socialid DROP COLUMN social_id;
ALTER TABLE socialid ADD COLUMN sns_id varchar(50);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE namecard DROP COLUMN name;
ALTER TABLE socialid DROP COLUMN sns_id;
