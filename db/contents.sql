drop table contents cascade;
create table contents (
	id		serial8	PRIMARY KEY,		
	mime		varchar	(50),		
	title		varchar	(100),		
	linkUrl		varchar	(100),		
	width		int2	,		
	height		int2	,		
	data		bytea	,		
	size		int	,		
	lastUpdated	timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
