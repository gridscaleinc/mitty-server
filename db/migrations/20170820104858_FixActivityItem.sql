
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE activity_item ADD COLUMN status varchar(20);
ALTER TABLE island ADD COLUMN geohash_l8 int8;
ALTER TABLE island ADD COLUMN geohash_l10 int8;
ALTER TABLE island ADD COLUMN geohash_l12 int8;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE activity_item DROP COLUMN status;
ALTER TABLE island DROP COLUMN geohash_l8;
ALTER TABLE island DROP COLUMN geohash_l10;
ALTER TABLE island DROP COLUMN geohash_l12;
