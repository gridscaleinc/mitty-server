
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

create table proposal (
  id                    serial8	PRIMARY KEY,
  reply_to_request_id   int8,
  contact_tel           varchar(20),
  contact_email         varchar(50),
  proposed_island_id    int8,
  proposed_island_id2   int8,
  gallery_id            int8,
  priceName1            varchar(50),
  price1                int,
  priceName2            varchar(50),
  price2                int,
  price_currency        varchar(3),
  price_info            varchar(1000),
  proposed_datetime1    timestamp,
  proposed_datetime2    timestamp,
  additional_info       text,
  proposer_id           int8,
  proposer_info         varchar(1000),
  proposed_datetime     timestamp,
  accept_status         varchar(20),
  accept_datetime       timestamp,
  confirm_tel           varchar(20),
  confirm_email         varchar(50),
  approval_status       varchar(20),
  approval_date         timestamp
);

create table footmark (
	id                 serial8 PRIMARY KEY,
	event_id           int8,
	island_id          int8,
	mitty_id           int,
	name_card_id       int8,
	picture_id         int8,
	seat_or_room_info  varchar(100),
	checkin_time       timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table proposal cascade;
drop table footmark cascade;
