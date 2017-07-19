
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE activity_item ADD COLUMN participation varchar (20);
ALTER TABLE activity_item ADD COLUMN calendar_url varchar (200);
ALTER TABLE request ADD COLUMN gallery_id int8;
ALTER TABLE events ADD COLUMN islandId2 int4;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE activity_item DROP COLUMN participation;
ALTER TABLE activity_item DROP COLUMN calendar_url;
ALTER TABLE request DROP COLUMN gallery_id;
ALTER TABLE events DROP COLUMN islandId2;
