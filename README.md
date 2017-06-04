# 開発

ローカルではdocker-compose startして開発を行う。

```
$ cd mitty-server
$ docker-compose start
$ docker-compose logs -f

```

# githubに反映

```
$ git push origin master
```

# サーバーに反映
```
$ ssh -i ~/.ssh/mitty-2017.pem ec2-user@52.196.151.53
$ cd go/src/mitty.co/mitty-server
$ git pull origin master
$ docker-compose logs -f
```

# Error

### input error
status: 400
```
{
  "errors":[
    "tag Required",
    "type Required"
  ]
}
```
### internal server error
status: 500
```
{
  "errors":[
    "database insert error"
  ]
}
```

# goose

```
# migrationファイル作成
$ goose -dir db/migrations create xxxx sql

# migrationファイルをDBに反映
$ goose -dir db/migrations postgres "postgres://root:mpNffadJrfnWpvZxnrZz2Zjz@dev.mitty.co:5432/mitty_db?sslmode=disable" up
```
