
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table activity (
  id	        	serial8	NOT NULL PRIMARY KEY,
  title		      varchar(200),
  main_event_id	int8,
  memo		      text,
  owner_id		  int8,
  created		    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated		    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table activity cascade;
