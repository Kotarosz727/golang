ハンドラ（Handler）:
ServeHTTPというメソッドを持ったインタフェースのこと

//インターフェースHTTPResponseWriterと構造体Requestへのポインタと言う２つの引数をとる
serveHTTP(http.ResponseWriter, *http.Request)