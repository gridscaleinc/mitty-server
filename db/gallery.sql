drop table gallery cascade ;
create table gallery (
	id		serial4	PRIMARY KEY,
	seq		int2,
	caption		varchar	(100),
	brief_info	text,
	content_id	int8,
	free_text	text
)
