package main

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	courses := []Course{
		{url: "https://vedio.speiyou.cn/63c3a33ce9194342a1b61828d9b92c93/2facd347620e476aaa1173cbe366df82-bcf3d32c2a27930d564f57799dd903d9-ld.mp4", name: "例题、挑战"},
		{url: "https://vedio.speiyou.cn/309874dc261c41708c2bee4173a40f07/124158a11511476ea672c33df6a3dab6-300c0976520b2b21c5449d233a8c717f-ld.mp4", name: "巩固、基础过关"},
		{url: "https://vedio.speiyou.cn/3e4d70712dae4b6e91f1a2f238840afe/b86c6aec8688484caf480c3a64b98a75-a4cf2a9dc50a019d86a0b9e45ff74b40-ld.mp4", name: "能力提升"},
	}
	for _, course := range courses {
		DownloadFileProgress(course.url, fmt.Sprintf("%s.mp4", course.name))
		fmt.Println(course.name, "download over")
	}
}
