
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table meeting (
        id			                                     serial8	    PRIMARY KEY,
        name                                        varchar(100),
        type			                                 varchar(20),
        created                timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated	               timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table meeting cascade;
