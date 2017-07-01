drop table contact cascade;
create table contact (
    id  serial8  PRIMARY KEY,
    mitty_id    int  ,
    name_card_id    int8,
    contacted_date    timestamp,
    related_event_id    int8
)