#!/bin/sh

id|grep postgres
if [ $? -ne 0 ] ; then
    echo ERROR: Not a dba user
    echo This is a dba user tool.
    echo Use \"su postgres\" to switch user
    exit
fi

echo create tables 

psql -d mitty -f activity.sql
psql -d mitty -f contents.sql
psql -d mitty -f events.sql
psql -d mitty -f gallery.sql
psql -d mitty -f island.sql
psql -d mitty -f users.sql
psql -d mitty -f vehicle.sql

psql -d mitty -c \\d+

