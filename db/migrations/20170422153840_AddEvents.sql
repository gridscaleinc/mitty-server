
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE if not exists events (
    id serial8 PRIMARY KEY,
    type VARCHAR (20),
    category VARCHAR (20),
    theme VARCHAR (50),
    tag VARCHAR (50),
    title VARCHAR (100),
    action text,
    start_datetime DATE,
    end_datetime DATE,
    allday_flag bool,
    islandId int4,
    logo_id int4,
    gallery_id int4,
    meeting_id int4,
    price_name1 VARCHAR (100),
    price1 integer,
    price_name2 VARCHAR (100),
    price2 integer,
    currency VARCHAR (3),
    price_info VARCHAR (200),
    description text,
    contact_tel VARCHAR (20),
    contact_fax VARCHAR (20),
    contact_mail VARCHAR (50),
    official_url VARCHAR (200),
    organizer VARCHAR (100),
    source_name VARCHAR (100),
    source_url VARCHAR (200),
    number_of_anticipants int4,
    anticipation VARCHAR (20),
    access_control VARCHAR (20),
    likes int4,
    status VARCHAR (20),
    language VARCHAR (10),
    created DATE,
    publisher_id int4,
    orgnization_id int4,
    lastupdated DATE,
    amender_id int4
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE events;
