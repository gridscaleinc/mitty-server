drop table contents cascade;
create table contents (
	id		serial8	PRIMARY KEY,		
	mime		varchar	(50),		
	name		varchar	(100),		
	link_url		varchar	(1000),		
	width		int2	,		
	height		int2	,		
	size		int	,		
	created    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated	timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
