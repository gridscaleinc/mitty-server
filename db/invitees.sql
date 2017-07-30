drop table invitees cascade;
create table invitees (
    id  serial8  PRIMARY KEY,
    invitation_id    int8  ,
    invitee_id  int,
    reply_status    varchar  (50),
    reply_time    timestamp  
);
