
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE conversation ALTER COLUMN speak_time TYPE timestamp USING '2017-07-05 00:00:00';

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE conversation ALTER COLUMN speak_time TYPE time USING '00:00:00';
