drop table users cascade;
create table users (
  id              serial PRIMARY KEY,
  name            varchar(30),
  user_name       varchar(30),
  password        varchar(255),
  access_token    varchar(255),
  mail_address    varchar(50),
  mail_confirmed  boolean,
  status          varchar	(20),
  icon             varchar(500),
  created         timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  updated         timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  UNIQUE          (user_name)
);
