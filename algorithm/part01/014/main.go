package main

import "fmt"

func main() {
	dominoes := [][]int{{1, 2}, {2, 1}, {1, 1}, {1, 2}, {2, 2}}
	result := numEquivDominoPairs1(dominoes)
	fmt.Println("result", result)
}

func numEquivDominoPairs1(dominoes [][]int) int {
	visitNum := 0
	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) ||
				(dominoes[i][1] == dominoes[j][0] && dominoes[i][0] == dominoes[j][1]) {
				visitNum++
			}
		}
	}
	return visitNum
}

func numEquivDominoPairs2(dominoes [][]int) int {
	visitNum := 0
	visitMap := make(map[string]int)
	for i := 0; i < len(dominoes); i++ {
		dm := dominoes[i]
		key := fmt.Sprintf("%d_%d", dm[0], dm[1])
		if dm[0] > dm[1] {
			key = fmt.Sprintf("%d_%d", dm[1], dm[0])
		}
		if v, ok := visitMap[key]; ok {
			visitMap[key] = v + 1
		} else {
			visitMap[key] = 1
		}
	}
	for _, v := range visitMap {
		visitNum = visitNum + v*(v-1)/2
	}
	return visitNum
}
