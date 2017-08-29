drop table offers cascade;
create table offers (
    id  serial8  PRIMARY KEY,
   from_mitty_id  int ,
   to_mitty_id int ,
   type  varchar (50),
   message  text ,
   reply_status varchar (50),
   offerred_id  int8 ,
   replied_id  int8 ,
   offerred_datetime  timestamp
);
