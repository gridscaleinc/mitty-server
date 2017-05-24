drop table request cascade;
create table request(
    id		int8	PRIMARY KEY,
    tag		varchar	(50),
    description		text	,
    for_activity_Id		int8,	
    preferred_datetime1 timestamp,	
    preferred_datetime2  timestamp	,
    preferred_price1		varchar(50),
    preferred_price2		varchar(50),
    start_place		varchar(100),
    terminate_place		varchar(100)
    oneway		bool	,
    status		varchar	(20)
    expiryDate		timestamp	,
    num_of_person		int2,	
    num_of_children		int2,	
    accepted_proposal_id		int8,	
    accepted_date		timestamp	,
    meeting_id		int8	,
    owner_id		int8	,
    created    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  	updated	timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);