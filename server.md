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
