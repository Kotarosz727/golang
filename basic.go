package main

import (
	//webサーバーを作成するためのライブラリ
	"net/http"
)

func sample() {
	//最もシンプルな構成
	//第一引数はネットワークアドレス、第二引数はハンドラ
	//ネットワークアドレスが「""」の場合、ポート80がデフォルトになる
	//ハンドラがnilの場合はデフォルトのマルチプレクサであるDefaultServeMuxが使われる
	http.ListenAndServe("", nil)
}