drop table activity_item cascade;
create table activity_item (
        activity_id                       int8 ,	
        event_id	                       int8	,
        title	                           varchar(200),
        memo		                   text,	
        notification                  bool,	
        notificationDateTime   date,	
        created    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated	timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY(activity_id, event_id)
)
