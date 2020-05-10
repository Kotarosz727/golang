ハンドラ（Handler）:
ServeHTTPというメソッドを持ったインタフェースのこと

//インターフェースHTTPResponseWriterと構造体Requestへのポインタと言う２つの引数をとる
serveHTTP(http.ResponseWriter, *http.Request)

//サーバー起動
go run XXX.go

//ドライバインストール
go get ~

//mysql ログイン（コンテナに入った後）
mysql -u root -p

//workbenchでコンテナ内のmysqlにアクセスできない場合
docker-compose down -v でボリュームを削除して、再度upするとアクセスできる