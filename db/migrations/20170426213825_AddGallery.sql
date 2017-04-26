
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table if not exists gallery (
  id		      serial4,
	seq		      int2,
	caption		  varchar	(100),
	brief_info	text,
	content_id	int8,
	free_text	  text,
	created     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated	    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(id,seq)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table gallery cascade;
