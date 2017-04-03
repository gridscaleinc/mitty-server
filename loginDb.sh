#!/bin/sh

CONTAINER_ID=$(docker ps |grep postgres |sed -e s/postgres.*$//)


docker ps 
echo "=========================================================="
echo Logining to PostgreSQL container id:  ${CONTAINER_ID}
echo 
echo Hint: using su postgres to switch user.
echo "     using psql to access database."
echo "     using \\h to get help."
echo "     using exit to get back to host os."
date
echo "=========================================================="
docker exec -it ${CONTAINER_ID} /bin/sh


