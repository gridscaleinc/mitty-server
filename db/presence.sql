drop table presence cascade;
create table presence (
    id  serial8  PRIMARY KEY,
    mitty_id    int  ,
    last_logedin    timestamp  ,
    online_status    varchar  (50),
    visiability    varchar  (50),
    traceablity    varchar  (50),
    latitude    float8  ,
    longitude    float8  ,
    located_time    timestamp  
    checkin_status    varchar  (50),
    checkin_time    timestamp  ,
    checkin_event_id    int  ,
    checkin_image_id    int8  
);