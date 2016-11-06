# mitty-server

# api format
レスポンスは全てjsonで行う。

# api authentication
認証はログインする度にサーバー側でaccess token発行クライアントに渡す。
クライアントは取得access tokenで認証付きリクエストおする。
