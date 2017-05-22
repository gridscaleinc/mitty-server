drop table meeting cascade;
create table meeting (
        id			                                     serial8	    PRIMARY KEY,
        name                                        varchar(100),
        type			                                 varchar(20),
        created                                     date 
)
