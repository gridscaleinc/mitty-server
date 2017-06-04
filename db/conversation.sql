drop table conversation cascade;
create table conversation (
        id			                                     serial8	    PRIMARY KEY,
        meeting_id                               int8,
        reply_to_id                               int8,
        speaking                                  text,
        speaker_id                               int8,
        speak_time                              date 
)
