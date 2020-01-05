package main

import "fmt"

func main() {
	result := minWindow1("aa", "aa")
	fmt.Println(result)
}

/**
初始，left指针和right指针都指向S的第一个元素.

将 right 指针右移，扩张窗口，直到得到一个可行窗口，亦即包含T的全部字母的窗口。

得到可行的窗口后，将left指针逐个右移，若得到的窗口依然可行，则更新最小窗口大小。

若窗口不再可行，则跳转至 2。
*/
func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	ss := []rune(s)
	ts := []rune(t)
	tm := make(map[rune]int)
	for _, e := range ts {
		if _, ok := tm[e]; ok {
			tm[e] = tm[e] + 1
		} else {
			tm[e] = 1
		}
	}
	required := len(tm)
	l, r, formed := 0, 0, 0
	ans := []int{-1, 0, 0}
	wm := make(map[rune]int)
	for r < len(ss) {
		c := ss[r]
		if _, ok := wm[c]; ok {
			wm[c] = wm[c] + 1
		} else {
			wm[c] = 1
		}
		if v, ok := tm[c]; ok {
			if wm[c] == v {
				formed++
			}
		}
		for l <= r && formed == required {
			c = ss[l]
			if ans[0] == -1 || r-l+1 < ans[0] {
				ans[0] = r - l + 1
				ans[1] = l
				ans[2] = r
			}
			wm[c] = wm[c] - 1
			if v, ok := tm[c]; ok {
				if wm[c] < v {
					formed--
				}
			}
			l++
		}
		r++
	}
	if ans[0] == -1 {
		return ""
	} else {
		return string(ss[ans[1] : ans[2]+1])
	}
}

/*
我们建立一个 filtered_S列表，其中包括 SS 中的全部字符以及它们在S的下标，但这些字符必须在 T中出现。

S = "ABCDDDDDDEEAFFBC" T = "ABC"
filtered_S = [(0, 'A'), (1, 'B'), (2, 'C'), (11, 'A'), (14, 'B'), (15, 'C')]

此处的(0, 'A')表示字符'A' 在字符串S的下表为0。
*/

type KeyPair struct {
	Key int
	Val rune
}

func minWindow1(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	ss := []rune(s)
	ts := []rune(t)
	tm := make(map[rune]int)
	for _, e := range ts {
		if _, ok := tm[e]; ok {
			tm[e] = tm[e] + 1
		} else {
			tm[e] = 1
		}
	}
	ms := make([]KeyPair, 0)
	for i, e := range ss {
		if _, ok := tm[e]; ok {
			ms = append(ms, KeyPair{Key: i, Val: e})
		}
	}
	required := len(tm)
	l, r, formed := 0, 0, 0
	ans := []int{-1, 0, 0}
	wm := make(map[rune]int)
	for r < len(ms) {
		c := ms[r].Val
		if _, ok := wm[c]; ok {
			wm[c] = wm[c] + 1
		} else {
			wm[c] = 1
		}
		if v, ok := tm[c]; ok {
			if wm[c] == v {
				formed++
			}
		}
		for l <= r && formed == required {
			c = ms[l].Val
			end := ms[r].Key
			start := ms[l].Key
			if ans[0] == -1 || end-start+1 < ans[0] {
				ans[0] = end - start + 1
				ans[1] = start
				ans[2] = end
			}
			wm[c] = wm[c] - 1
			if v, ok := tm[c]; ok {
				if wm[c] < v {
					formed--
				}
			}
			l++
		}
		r++
	}
	if ans[0] == -1 {
		return ""
	} else {
		return string(ss[ans[1] : ans[2]+1])
	}
}

func minWindow2(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hash := make([]int, 256)
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	l, count, max, results := 0, len(t), len(s)+1, ""
	for r := 0; r < len(s); r++ {
		hash[s[r]]--
		if hash[s[r]] >= 0 {
			count--
		}
		for l < r && hash[s[l]] < 0 {
			hash[s[l]]++
			l++
		}
		if count == 0 && max > r-l+1 {
			max = r - l + 1
			results = s[l : r+1]
		}
	}

	return results
}
