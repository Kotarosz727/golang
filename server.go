package main

import (
	"net/http"
)

func ain() {
	server := http.Server{
		Addr:	"127.0.0.1:8080",
		Handler:nil,
	}
	//HTTPS対応
	//cert.pemがSSL証明書,key.pemがサーバー用の秘密鍵
	server.ListenAndServeTLS("cert.pem", "key.pem")
}