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
cd mitty-server
./loginDb.sh

su postgres    -- DBA user 
cd /var/db     -- DB のワークエリア
psql -d mitty  -- mitty Database にログイン

---- psqlの基本コマンド
1)  \list    ----> database をリストアップ
2)  \d       ----> table をリストアップ
3)  \q       ----> psqlから抜ける。  exitじゃないよ。
4)  \h       ----> Help
5   select 、insert .....

# api format
レスポンスは全てjsonで行う。

# api authentication
認証はログインする度にサーバー側でaccess token発行クライアントに渡す。
クライアントは取得access tokenで認証付きリクエストおする。

# server log
```
$ ssh mitty-server
$ cd go/src/mitty.co/mitty-server
$ docker-compose logs -f
```
