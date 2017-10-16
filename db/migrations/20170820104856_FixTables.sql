-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE contact ADD COLUMN meeting_id int8;
ALTER TABLE conversation ADD COLUMN type varchar(50);
ALTER TABLE conversation ADD COLUMN latitude float8;
ALTER TABLE conversation ADD COLUMN longitude float8;
ALTER TABLE socialLink ADD COLUMN contact_id int8;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE contact DROP COLUMN meeting_id;
ALTER TABLE conversation DROP COLUMN type;
ALTER TABLE conversation DROP COLUMN latitude;
ALTER TABLE conversation DROP COLUMN longitude;
ALTER TABLE socialLink DROP COLUMN contact_id;
