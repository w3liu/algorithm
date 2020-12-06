## Algorithm
### 1. 题目
```
1162. 地图分析
```
### 2. 题目描述
```
你现在手里有一份大小为 N x N 的 网格 grid，上面的每个 单元格 都用 0 和 1 标记好了。其中 0 代表海洋，1 代表陆地，请你找出一个海洋单元格，这个海洋单元格到离它最近的陆地单元格的距离是最大的。

我们这里说的距离是「曼哈顿距离」（ Manhattan Distance）：(x0, y0) 和 (x1, y1) 这两个单元格之间的距离是 |x0 - x1| + |y0 - y1| 。

如果网格上只有陆地或者海洋，请返回 -1。

 

示例 1：
```
[!qr](./images/1206_a_1.jpeg)
```

输入：[[1,0,1],[0,0,0],[1,0,1]]
输出：2
解释： 
海洋单元格 (1, 1) 和所有陆地单元格之间的距离都达到最大，最大距离为 2。
示例 2：
```
[!qr](./images/1206_a_2.jpeg)
```

输入：[[1,0,0],[0,0,0],[0,0,0]]
输出：4
解释： 
海洋单元格 (2, 2) 和所有陆地单元格之间的距离都达到最大，最大距离为 4。

```

### 3. 解答：
```golang
func maxDistance(grid [][]int) int {
	var dx = []int{0, 0, 1, -1}
	var dy = []int{1, -1, 0, 0}
	queue := list.New()
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue.PushBack([]int{i, j})
			}
		}
	}
	hasOcean := false
	var point []int
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		point = ele.Value.([]int)
		x := point[0]
		y := point[1]
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1
			hasOcean = true
			queue.PushBack([]int{newX, newY})
		}
	}
	if point == nil || !hasOcean {
		return -1
	}
	return grid[point[0]][point[1]] - 1
}
```
### 4. 说明

相信对于Tree的BFS大家都已经轻车熟路了：要把root节点先入队，然后再一层一层的无脑遍历就行了。

对于图的BFS也是一样滴～ 与Tree的BFS区别如下：
1、tree只有1个root，而图可以有多个源点，所以首先需要把多个源点都入队。
2、tree是有向的因此不需要标志是否访问过，而对于无向图来说，必须得标志是否访问过！
并且为了防止某个节点多次入队，需要在入队之前就将其设置成已访问！

这是一道典型的BFS基础应用，为什么这么说呢？
因为我们只要先把所有的陆地都入队，然后从各个陆地同时开始一层一层的向海洋扩散，那么最后扩散到的海洋就是最远的海洋！
并且这个海洋肯定是被离他最近的陆地给扩散到的！
下面是扩散的图示，1表示陆地，0表示海洋。每次扩散的时候会标记相邻的4个位置的海洋：

你可以想象成你从每个陆地上派了很多支船去踏上伟大航道，踏遍所有的海洋。每当船到了新的海洋，就会分裂成4条新的船，向新的未知海洋前进（访问过的海洋就不去了）。如果船到达了某个未访问过的海洋，那他们是第一个到这片海洋的。很明显，这么多船最后访问到的海洋，肯定是离陆地最远的海洋。
