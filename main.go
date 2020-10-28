package main

import (
	"go-network-disk/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/download", handler.DownloadHandler)

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Printf("http.ListenAndServe fail,err: %v \n", err)
		return
	}
}
