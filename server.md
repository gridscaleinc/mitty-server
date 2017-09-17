# aws
aws consoleで書くサーバを起動する。

# app
nginx
```
$ sudo yum install nginx
$ sudo /etc/init.d/nginx start
$ curl localhost
```
go
```
$ mkdir dl
$ cd dl
$ wget https://storage.googleapis.com/golang/go1.9.linux-amd64.tar.gz
$ sudo tar -C /usr/local -xzf go1.9.linux-amd64.tar.gz
$ vim ~/.bashrc
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
$ go version
go version go1.9 linux/amd64
$
```

# db
postgres
```
$ sudo yum install -y postgresql95 postgresql95-server postgresql95-libs postgresql95-contrib
$ sudo service postgresql95 initdb 
$ sudo /etc/init.d/postgresql95 start
$ sudo su - postgres
$ createdb mitty_db
$ psql

CREATE ROLE mitty_user LOGIN CREATEDB SUPERUSER PASSWORD 'PxeFKA9nXawKhyrFCi2Ajrenyzkocy';

$ sudo vim /var/lib/pgsql95/data/pg_hba.conf 
local   all             all                                    trust
host    all             all             0.0.0.0/0              md5
$ sudo vim /var/lib/pgsql95/data/postgresql.conf
listen_addresses = '*'

$ sudo /etc/init.d/postgresql95 restart
```

# elasticsearch

install
```
$ sudo yum -y install java-1.8.0-openjdk
$ sudo alternatives --config java
select number

$ sudo rpm -i https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-5.3.2.rpm
```

start
```
$ sudo service elasticsearch start
Starting elasticsearch: OpenJDK 64-Bit Server VM warning: INFO: os::commit_memory(0x0000000085330000, 2060255232, 0) failed; error='Cannot allocate memory' (errno=12)

$ sudo vim /etc/init.d/elasticsearch
export ES_JAVA_OPTS="-Xms512m -Xmx512m"

$ sudo service elasticsearch start
```
config
```
$ sudo chkconfig --add elasticsearch

$ cd /usr/share/elasticsearch/
$ sudo bin/elasticsearch-plugin install analysis-kuromoji

$ sudo service elasticsearch start
```

setting
```
$ curl localhost:9200/_cluster/health?pretty
{
  "cluster_name" : "elasticsearch",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 13,
  "active_shards" : 13,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 100.0
}
```
