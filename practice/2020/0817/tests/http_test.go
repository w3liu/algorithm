package tests

import "testing"

func TestDoRequest(t *testing.T) {
	url := "http://www.baidu.com"
	if err := DoRequest(url); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDoRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "http://www.baidu.com"
		if err := DoRequest(url); err != nil {
			b.Fatal(err)
		}
	}
}
