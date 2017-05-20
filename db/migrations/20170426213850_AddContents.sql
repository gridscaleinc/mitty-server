
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists contents (
  id		      serial8	PRIMARY KEY,
	mime		    varchar	(50),
	name		    varchar	(100),
	link_url		varchar	(1000),
	width		    int2,
	height		  int2,
	size		    int,
  owner_id    int,
	created     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated	    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table contents cascade;
