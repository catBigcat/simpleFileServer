package main

import (
	"net/http"
	"os"
)

func main() {
	add := os.Getenv("FILE_SERVER_ADDR")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.Handle("/files/", handlerMd(http.StripPrefix("/files/", http.FileServer(http.Dir("./files")))))
	err := http.ListenAndServe(add, nil)
	if err != nil {
		println(err.Error())
	}
}
