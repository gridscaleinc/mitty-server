drop table invitation cascade;
create table invitation (
    id  serial8  PRIMARY KEY,
    invitater_id    int  ,
    for_type    varchar  (50),
    id_of_type    varchar  (50),
    message    text  ,
    time_of_invitation    timestamp 
);
