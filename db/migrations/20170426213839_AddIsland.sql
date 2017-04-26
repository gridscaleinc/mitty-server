
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists island (
  id              serial PRIMARY KEY,
  nickname        varchar(89),
  name            varchar(80),
  logo_id         int4,
  description     text,
  category        varchar(20),
  mobility_type   varchar(20),
  reality_type    varchar(20),
  ownership_type  varchar(20),
  owner_name      varchar(80),
  owner_id        int4,
  creator_id      int4,
  meeting_id      int4,
  gallery_id      int4,
  tel             varchar(20),
  fax             varchar(20),
  mailaddress     varchar(50),
  webpage         varchar(50),
  likes           int,
  country_code    varchar(2),
  country_name    varchar(30),
  state           varchar(30),
  city            varchar(30),
  postcode        varchar(20),
  thoroghfare     varchar(30),
  subthroghfare   varchar(30),
  building_name   varchar(50),
  floor_number    varchar(3),
  room_number     varchar(10),
  address1        varchar(100),
  address2        varchar(100),
  address3        varchar(100),
  latitude        float8,
  longitude       float8,
  created         timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated	        timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table island cascade;
