
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table conversation (
        id			                                     serial8	    PRIMARY KEY,
        meeting_id                               int8,
        reply_to_id                               int8,
        speaking                                  text,
        speaker_id                               int8,
        speak_time                              date
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table conversation cascade;
