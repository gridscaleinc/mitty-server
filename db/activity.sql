drop table activity cascade;
create table activity (
       id	        	           serial8	NOT NULL PRIMARY KEY,
       itle		                   varchar(200),
       main_event_id		int8,
      memo		           text,
      owner_id		        int8,
      created		           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated		           timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
