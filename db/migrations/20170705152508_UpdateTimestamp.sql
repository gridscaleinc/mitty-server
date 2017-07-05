
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE activity_item ALTER COLUMN notificationDateTime TYPE timestamp;
ALTER TABLE conversation ALTER COLUMN speak_time TYPE time USING '00:00:00';
ALTER TABLE events ALTER COLUMN start_datetime TYPE timestamp;
ALTER TABLE events ALTER COLUMN end_datetime TYPE timestamp;
ALTER TABLE events ALTER COLUMN lastupdated TYPE timestamp;
ALTER TABLE request ALTER COLUMN preferred_datetime1 TYPE timestamp;
ALTER TABLE request ALTER COLUMN preferred_datetime2 TYPE timestamp;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE activity_item ALTER COLUMN notificationDateTime TYPE date;
ALTER TABLE conversation ALTER COLUMN speak_time TYPE date;
ALTER TABLE events ALTER COLUMN start_datetime TYPE date;
ALTER TABLE events ALTER COLUMN end_datetime TYPE date;
ALTER TABLE events ALTER COLUMN lastupdated TYPE date;
ALTER TABLE request ALTER COLUMN preferred_datetime1 TYPE date;
ALTER TABLE request ALTER COLUMN preferred_datetime2 TYPE date;
