drop table events cascade;
create table events (
        id			                                     serial8	    PRIMARY KEY,
        type			                                 varchar(20),
        category		                             varchar(20),
        theme			                             varchar(50),
        title			                                 varchar(100),
        action		                                 text,
        start_datetime	 	                     date,
        end_datetime		                     date,
        allday_flag		                         bool,
        islandId		                             int4,
        logo_id                                     int4,
        gallery_id                                 int4,
        meeting_id                               int4,
        price_name1                            varchar(100),
        price1                                       money,
        price_name2                            varchar(100),
        price2                                       money,
        currency                                   varchar(3),
        price_info                                 varchar(200),
        description                               text,
        contact_tel                               varchar(20),
        contact_fax                              varchar(20),
        contact_mail                            varchar(50),
        official_url                                varchar(200),
        organizer                                 varchar(100),
        source_name                          varchar(100),
        source_url                               varchar(200),
        number_of_anticipants            int4,
        anticipation                              varchar(20),
        access_control                        varchar(20),
        likes                                         int4,
        status                                       varchar(20),
        language                                  varchar(10),
        created                                     date,
        publisher_id                             int4,
        orgnization_id                          int4,
        lastupdated                              date,
        amender_id                             int4
        );
