drop table gallery cascade ;
create table gallery (
	id		serial4	PRIMARY KEY,
	seq		int2,
	caption		varchar	(100),
	briefInfo	text,
	url		varchar	(100),
	contentId	int8,
	freeText	text
)
