package main

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	courses := []Course{
		{url: "https://vedio.speiyou.cn/63c3a33ce9194342a1b61828d9b92c93/2facd347620e476aaa1173cbe366df82-bcf3d32c2a27930d564f57799dd903d9-ld.mp4", name: "第9讲"},
	}
	for _, course := range courses {
		DownloadFileProgress(course.url, fmt.Sprintf("%s.mp4", course.name))
		fmt.Println(course.name, "download over")
	}
}
