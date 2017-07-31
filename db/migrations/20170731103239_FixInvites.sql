
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
drop table invitation cascade;
create table if not exists invitation (
    id  serial8         PRIMARY KEY,
    invitater_id        int,
    for_type            varchar(50),
    id_of_type          varchar(50),
    message             text,
    time_of_invitation  timestamp
);
create table invitees (
    id  serial8         PRIMARY KEY,
    invitation_id       int8,
    invitee_id          int,
    reply_status        varchar(50),
    reply_time          timestamp
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
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
drop table invitation cascade;
drop table invitees cascade;
