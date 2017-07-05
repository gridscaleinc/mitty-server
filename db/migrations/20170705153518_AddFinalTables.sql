
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists profile (
  id                  serial8 PRIMARY KEY,
  mitty_id            int,
  gender              varchar(10),
  one_word_speech     varchar(200),
  constellation       varchar(20),
  home_island_id      int8,
  birth_island_id     int8,
  age_group           varchar(20),
  appearance_tag      varchar(20),
  occupation_tag1     varchar(20),
  occupation_tag2     varchar(20),
  occupation_tag3     varchar(20),
  hobby_tag1          varchar(20),
  hobby_tag2          varchar(20),
  hobby_tag3          varchar(20),
  hobby_tag4          varchar(20),
  hobby_tag5          varchar(20)
);

create table if not exists namecard (
    id                 serial8 PRIMARY KEY,
    mitty_id           int,
    business_name      varchar(200),
    business_sub_name  varchar(200),
    business_title     varchar(200),
    address_line1      varchar(100),
    address_line2      varchar(100),
    phone              varchar(20),
    fax                varchar(20),
    mobile_phone       varchar(20),
    webpage            varchar(100),
    email              varchar(100),
    created            timestamp    not null  DEFAULT CURRENT_TIMESTAMP,
    updated            timestamp    not null  DEFAULT CURRENT_TIMESTAMP
);

create table if not exists contact (
    id                  serial8  PRIMARY KEY,
    mitty_id            int,
    name_card_id        int8,
    contacted_date      timestamp,
    related_event_id    int8
);

create table if not exists presence (
    id                  serial8 PRIMARY KEY,
    mitty_id            int,
    last_logedin        timestamp,
    online_status       varchar(50),
    visiability         varchar(50),
    traceablity         varchar(50),
    latitude            float8,
    longitude           float8,
    located_time        timestamp,
    checkin_status      varchar(50),
    checkin_time        timestamp,
    checkin_event_id    int,
    checkin_image_id    int8
);

create table if not exists invitation (
    id                  serial8 PRIMARY KEY,
    invitator_id        int,
    for_type            varchar(50),
    id_of_type          varchar(50),
    message             text,
    time_of_invitation  timestamp,
    reply_status        varchar(50),
    reply_time          timestamp
);

create table if not exists socialId (
   id                   serial8 PRIMARY KEY,
   contactee_id         int,
   realm                varchar(50),
   social_id            int8
);

create table if not exists socialLink (
    id	                serial8 PRIMARY KEY,
    mitty_id		        int,
    social_id		        int8
);

create table if not exists likes (
    id                  serial8 PRIMARY KEY,
    mitty_id            int,
    entity_type         varchar(50),
    entity_id           int8
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table profile cascade;
drop table namecard cascade;
drop table contact cascade;
drop table presence cascade;
drop table invitation cascade;
drop table socialId cascade;
drop table socialLink cascade;
drop table likes cascade;
