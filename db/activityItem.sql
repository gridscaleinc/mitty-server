drop table activity_item cascade;
create table activity_item (
        id	        	                       serial8	NOT NULL PRIMARY KEY,
        activity_id                        int8 ,
        event_id	                        int8	,
        title	                                varchar(200),
        memo		                       text,
        notification                      bool,
        notificationdatetime        timestamp,
        participation		               varchar	(20),
        calendar_url		               varchar	(200),
        status                         varchar (20),
        created                           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated	                       timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
)
