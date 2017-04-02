# mitty-server

# install docker for mac
https://docs.docker.com/docker-for-mac/

# start web server
```
$ git clone git@github.com:gridscaleinc/mitty-server.git
$ cd mitty-server
$ docker network create --driver bridge mitty_docker
$ docker-compose up -d
```
http://localhost:8000


# api format
レスポンスは全てjsonで行う。

# api authentication
認証はログインする度にサーバー側でaccess token発行クライアントに渡す。
クライアントは取得access tokenで認証付きリクエストおする。
