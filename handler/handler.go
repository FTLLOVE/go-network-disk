package handler

import (
	"encoding/json"
	"go-network-disk/meta"
	"go-network-disk/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// uploadHandler: 上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			log.Printf("ioutil.ReadFile fail,err: %v \n", err)
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == http.MethodPost {
		// 接收文件流存储到本地目录
		file, header, err := r.FormFile("file")
		if err != nil {
			log.Printf("r.FormFile fail,err: %v \n", err.Error())
			return
		}
		defer file.Close()
		// 创建新文件
		fileMeta := &meta.FileMeta{
			FileName: header.Filename,
			Location: "./tmp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("os.Create fail,err: %v \n", err.Error())
			return
		}
		defer newFile.Close()
		// 拷贝到新文件
		fileMeta.FileSize, err = io.Copy(newFile, file)
		if err != nil {
			log.Printf("io.Copy fail,err: %v \n", err.Error())
			return
		}
		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	} else {

	}
}

// UploadSuccessHandler: 上传文件成功的提示
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload success")
}

// GetFileMetaHandler: 获取文件元信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(s)
	data, err := json.Marshal(fileMeta)
	if err != nil {
		log.Printf("json.Marshal fail,err: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// DownloadHandler: 下载文件元
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	s := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(s)
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		log.Printf("os.Open fail,err: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("ioutil.ReadAll fail,err: %v \n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octect-stream")
	w.Header().Set("Content-disposition", "attachement;filename=\""+fileMeta.FileName+"\"")
	w.Write(data)
}
