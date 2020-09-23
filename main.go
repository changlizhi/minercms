package main

import (
	"log"
	"net/http"
)

func CmsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("123"))
	return
}
func StartAPI() {
	http.HandleFunc("/cms", CmsHandler)
	err := http.ListenAndServe("localhost:8888", nil)
	if err != nil {
		log.Fatal("服务端报错：", err.Error())
	}
}
func main() {
	StartAPI()
}
