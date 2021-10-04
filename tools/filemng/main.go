package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

const FilesPath = "files"
const ImagesPath = "images"
const Port = ":9219"

func main() {

	http.HandleFunc("/upload/file", handleUploadFile)
	http.HandleFunc("/upload/image", handleUploadImage)
	http.HandleFunc("/download", handleDownload)

	server := &http.Server{
		Addr:    Port,
		Handler: nil,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Server run fail" + err.Error())
		}
	}()

	log.Println("file server starting···")
	time.Sleep(1 * time.Second)

	log.Println("file server started")
	log.Println("listen at " + Port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL)

	select {
	// wait on kill signal
	case <-ch:
		err := server.Close()
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("file server stopped")
		}
	}
}

func handleUploadFile(w http.ResponseWriter, request *http.Request) {
	handleUpload(w, request, FilesPath)
}

func handleUploadImage(w http.ResponseWriter, request *http.Request) {
	handleUpload(w, request, ImagesPath)
}

func handleUpload(w http.ResponseWriter, request *http.Request, basePath string) {
	//文件上传只允许POST方法
	if request.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("error|Method not allowed"))
		return
	}

	//从表单中读取文件
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		_, _ = io.WriteString(w, "error|Read file error")
		return
	}
	//defer 结束时关闭文件
	defer file.Close()

	var fid = strings.ReplaceAll(uuid.New().String(), "-", "")
	var ext = path.Ext(fileHeader.Filename)
	var newName = fmt.Sprintf("%s%s", fid, ext)

	log.Println("filename: "+fileHeader.Filename, "newName:"+newName)

	_, err = os.Stat(basePath)
	if err != nil && !os.IsNotExist(err) {
		_, _ = io.WriteString(w, "error|os.Stat error")
		return
	}
	_ = os.Mkdir(basePath, 0644)
	//创建文件
	newFile, err := os.Create(basePath + "/" + newName)
	if err != nil {
		_, _ = io.WriteString(w, "error|Create file error")
		return
	}
	//defer 结束时关闭文件
	defer newFile.Close()

	//将文件写到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		_, _ = io.WriteString(w, "error|Write file error")
		return
	}
	_, _ = io.WriteString(w, fmt.Sprintf("%s|%s", "success", newName))
}

func handleDownload(w http.ResponseWriter, request *http.Request) {
	//文件下载只允许GET方法
	if request.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("Method not allowed"))
		return
	}
	//文件名
	filename := request.URL.Query().Get("fileName")
	if filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	log.Println("filename: " + filename)
	//打开文件
	file, err := os.Open(FilesPath + "/" + filename)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
	//结束后关闭文件
	defer file.Close()

	//设置响应的header头
	w.Header().Add("Content-type", "application/octet-stream")
	w.Header().Add("content-disposition", "attachment; filename=\""+filename+"\"")
	//将文件写至responseBody
	_, err = io.Copy(w, file)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, "Bad request")
		return
	}
}
