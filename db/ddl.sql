create table users (
  id              serial PRIMARY KEY,
  username        varchar(255) NOT NULL,
  password        varchar(255) NOT NULL,
  access_token    varchar(255) NOT NULL,
  name            varchar(255),
  created         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (username)
);
