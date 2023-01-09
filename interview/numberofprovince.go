package main

func main() {

}

func findCircleNumDFS(isConnected [][]int) int {

	n, count := len(isConnected), 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		count++
		dfs(isConnected, visited)
	}
	return count
}

func dfs(isConnected [][]int, row int, visited []bool) {
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] && isConnected[row][i] == 1 {
			visited[i] = true
			dfs(isConnected, i, visited)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	// create disjoint union set
	// 1...n, union find parent, if not existm increase count and add to set
	// if exist, continue
	n := len(isConnected)
	set := &DisjointUnionSet{
		parent: make([]int, n),
	}
	for i := range set.parent {
		set.parent[i] = -1
	}

	for i := range isConnected {
		for j := range isConnected[i] {
			if i == j {
				continue
			}
			if isConnected[i][j] == 1 {
				set.Union(i, j)
			}
		}
	}

	count := 0
	mapping := make(map[int]int)
	for i := 0; i < n; i++ {
		pi := set.Find(i)
		if _, ok := mapping[pi]; ok {
			continue
		}
		count++
		mapping[pi]++
	}

	return count
}

type DisjointUnionSet struct {
	parent []int
}

func (dsu *DisjointUnionSet) Find(x int) int {
	if dsu.parent[x] == -1 {
		return x
	}
	return dsu.Find(dsu.parent[x])
}

func (dsu *DisjointUnionSet) Union(x, y int) bool {
	px, py := dsu.Find(x), dsu.Find(y)
	if px == py {
		return false
	}
	dsu.parent[py] = px
	return true
}
