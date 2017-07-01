drop table namecard cascade;
create table namecard (
    id    serial8  PRIMARY KEY,    
    mitty_id    int      
    business_name    varchar(200),    
    business_sub_name    varchar(200),  
    business_title    varchar(200),  
    address_line1    varchar(100),    
    address_line2    varchar(100),    
    phone    varchar(20),    
    fax    varchar(20),    
    mobile_phone  varchar(20),    
    webpage    varchar(100),    
    email    varchar(100),    
    created    timestamp    not null  DEFAULT CURRENT_TIMESTAMP,
    updated    timestamp    not null  DEFAULT CURRENT_TIMESTAMP,
);