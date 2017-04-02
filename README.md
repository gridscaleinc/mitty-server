# mitty-server

# install docker for mac
https://docs.docker.com/docker-for-mac/

# mitty-serverチェックアウト
```
$ git clone git@github.com:gridscaleinc/mitty-server.git
$ cd mitty-server
$ docker network create --driver bridge mitty_docker

# mitty-server起動
$ cd mitty-server
$ docker-compose up -d
```

#起動確認
http://localhost:8000/api/status

下記JSON内容が表示されれば、OK
{"name":"mitty api server","version":"0.0.1"}

# docker commands

```
$ docker images
$ docker ps
$ docker-compose start
$ docker-compose stop
```

##コンテーナーの中にはいで作業する。(postgresを例として）
# container idを確認する

$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                         NAMES
314874e143c5        postgres:9.6.0      "/docker-entrypoint.s"   3 minutes ago       Up 3 minutes        0.0.0.0:5432->5432/tcp        mitty_postgres
4796ea8c4e65        nginx:1.11.5        "nginx -g 'daemon off"   3 minutes ago       Up 3 minutes        0.0.0.0:80->80/tcp, 443/tcp   mitty_nginx

$ docker exec -it 314874e143c5 /bin/bash
root@314874e143c5:/# su postgres
$ psql
psql (9.6.0)
Type "help" for help.

postgres=# 
```

# api format
レスポンスは全てjsonで行う。

# api authentication
認証はログインする度にサーバー側でaccess token発行クライアントに渡す。
クライアントは取得access tokenで認証付きリクエストおする。
