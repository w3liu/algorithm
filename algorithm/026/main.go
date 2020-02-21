package main

import (
	"fmt"
)

func main() {

	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 4}, {4, 5}, {4, 6}, {6, 7}}
	ans := findMinHeightTrees(8, edges)
	for _, item := range ans {
		fmt.Println(item)
	}
}

func findMinHeightTrees(n int, edges [][]int) []int {
	// 结果集
	var ans = make([]int, 0)
	if n == 1 {
		ans = append(ans, 0)
		return ans
	}

	// 记录节点的度
	var degree = make([]int, n)
	// 邻接矩阵
	var graph = make([][]int, n)
	// 存放叶子节点的队列
	var queue = make(map[int]struct{}, 0)

	for _, item := range edges {
		first := item[0]
		second := item[1]
		// 设置节点的度
		degree[first]++
		degree[second]++
		// 构建邻接表
		graph[first] = append(graph[first], second)
		graph[second] = append(graph[second], first)
	}

	// 度为1的为叶子节点，将其放入队列中
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue[i] = struct{}{}
		}
	}

	var lst []int

	// 遍历叶子节点
	for len(queue) != 0 {
		ans = make([]int, 0)
		q := make(map[int]struct{})
		for k := range queue {
			// 将叶子节点放入结果集中
			ans = append(ans, k)
			// 获取叶子节点的邻接表
			lst = graph[k]
			for j := 0; j < len(lst); j++ {
				node := lst[j]
				// 更新节点的度
				degree[node] = degree[node] - 1
				// 如果度为1了，说明变成了新的叶子节点，则将其压入队列
				if degree[node] == 1 {
					q[node] = struct{}{}
				}
			}
		}
		queue = q
	}

	return ans
}
