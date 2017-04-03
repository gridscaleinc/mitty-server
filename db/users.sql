drop table users cascade;
create table users (
  id                  	serial PRIMARY KEY,
  name			varchar(30),	
  nickName		varchar(30),
  password		varchar(255),		
  accessToken		varchar(255),		
  mailAddress		varchar(50),		
  mailConfirmed		boolean,			
  status		varchar	(20),		
  created		timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  updated		timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (nickName)
);
