drop table gallery cascade ;
create table gallery (
	id		serial4	,
	seq		int2,
	caption		varchar	(100),
	brief_info	text,
	content_id	int8,
	free_text	text,
	created    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated	timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(id,seq)
)
