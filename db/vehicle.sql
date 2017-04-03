drop table vehicle cascade;
create table vehicle(
    islandId              int4 PRIMARY KEY,
    serviceNo             varchar(64),
    category              varchar(20),
    tonage                int2,
    numberOfSeats         int2,
    apartPlace            varchar(100),
    terminalPlace         varchar(100)
);
