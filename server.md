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
