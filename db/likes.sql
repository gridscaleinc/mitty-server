drop table likes cascade;
create table likes (
    id    serial8  PRIMARY KEY,
    mitty_id    int  
    entity_type    varchar(50),
    entity_id    int8  
);