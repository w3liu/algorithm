package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	ans := findAnagrams(s, p)
	fmt.Println(ans)
}

func findAnagrams(s, p string) (ans []int) {
	sl := len(s)
	pl := len(p)
	ans = make([]int, 0)
	if pl > sl {
		return
	}
	var scnt, pcnt [26]int
	for i, v := range p {
		scnt[s[i]-'a']++
		pcnt[v-'a']++
	}

	if scnt == pcnt {
		ans = append(ans, 0)
	}

	for i, v := range s[:sl-pl] {
		scnt[v-'a']--
		scnt[s[i+pl]-'a']++
		if scnt == pcnt {
			ans = append(ans, i+1)
		}
	}
	return
}
