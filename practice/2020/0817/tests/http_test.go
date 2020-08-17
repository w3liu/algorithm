package tests

import "testing"

func TestDoRequest(t *testing.T) {
	url := "http://www.baidu.com"
	if err := DoRequest(url); err != nil {
		t.Fatal(err)
	}
}
