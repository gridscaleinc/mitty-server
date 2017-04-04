
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists users (
  id              serial PRIMARY KEY,
  name            varchar(30),
  user_name       varchar(30),
  password        varchar(255),
  access_token    varchar(255),
  mail_address    varchar(50),
  mail_confirmed  boolean,
  status          varchar	(20),
  created         timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  updated         timestamp  not null DEFAULT CURRENT_TIMESTAMP,
  UNIQUE          (user_name)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users cascade;
