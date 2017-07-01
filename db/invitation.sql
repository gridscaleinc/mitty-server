drop table invitation cascade;
create table invitation (
    id  serial8  PRIMARY KEY,
    invitator_id    int  ,
    for_type    varchar  (50),
    id_of_type    varchar  (50),
    message    text  ,
    time_of_invitation    timestamp  ,
    reply_status    varchar  (50),
    reply_time    timestamp  
);

