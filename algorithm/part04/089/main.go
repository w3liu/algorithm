package main

func main() {

}

func groupAnagrams(strs []string) [][]string {
	var cm = make(map[[26]int][]string)
	for _, str := range strs {
		ele := [26]int{}
		for _, s := range str {
			ele[s-'a']++
		}
		cm[ele] = append(cm[ele], str)
	}
	var result = make([][]string, 0, len(strs))
	for _, v := range cm {
		result = append(result, v)
	}
	return result
}
