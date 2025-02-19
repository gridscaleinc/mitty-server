drop table request cascade ;
create table request (
        id			                                     serial8	    PRIMARY KEY,
        title			                                 varchar(100),
        tag			                                 varchar(20),
        description                                text,
        for_activity_id                           int8,
        preferred_datetime1                 timestamp,
        preferred_datetime2                 timestamp,
        preferred_price1                       integer,
        preferred_price2                       integer,
        currency                                    varchar(3),
        start_place                                varchar(200),
        terminate_place                        varchar(200),
        oneway                                     bool,
        status                                        varchar(20),
        expiry_date                                 date,
        num_of_person                          integer,
        num_of_children                        integer,
        gallery_id		                              int8	,
        accepted_proposal_id               int8,
        accepted_date                           timestamp,
        meeting_id                                 int8,
        owner_id                                    int8,
	    created    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	    updated   timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)
;