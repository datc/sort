package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("ASTAR\t\t")
	display(data)
	dfs(1, 1)
}

var (
	data = [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 9, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	}
	visited [][]bool
	target  = 9
	m       = len(data)
	n       = len(data[0])
	start   = time.Now()
)

func init() {
	if m <= 0 {
		return
	}
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
}

/**
display map
 */
func display(data [][]int) {
	for _, it := range data {
		for _, itt := range it {
			fmt.Print(itt, "  ")
		}
		fmt.Println()
	}
}

/**
print the current rute
 */
func rute(r [][]bool) {
	for _, it := range r {
		for _, itt := range it {
			fmt.Print(itt, "  ")
		}
		fmt.Println()
	}
}

/**
dfs algorithm
 */
func dfs(i, j int) {
	v := data[i][j]
	fmt.Println(fmt.Sprintf("(%d,%d):%d\t\t", i, j, v))
	display(data)
	if v == target {
		fmt.Println("target is found, costs ", time.Since(start))
		os.Exit(0)
	}
	if v != 0 {
		return
	}
	visited[i][j] = true
	data[i][j] = 2
	for _, it := range next(i, j) {
		dfs(it.i, it.j)
	}
	visited[i][j] = false
	data[i][j] = 0
}

/**
visit step
 */
type vst struct {
	i, j int
}

/**
return all the posibile rute of (i,j): up,down,left and right
 */
func next(i, j int) []vst {
	vsts := make([]vst, 0, 4)
	if safe(i-1, j) {
		vsts = append(vsts, vst{i: i - 1, j: j})
	}
	if safe(i+1, j) {
		vsts = append(vsts, vst{i: i + 1, j: j})
	}
	if safe(i, j-1) {
		vsts = append(vsts, vst{i: i, j: j - 1})
	}
	if safe(i, j+1) {
		vsts = append(vsts, vst{i: i, j: j + 1})
	}
	return vsts
}

/**
check the next step if safe: out of range or visited
 */
func safe(i, j int) bool {
	if i >= m || j >= n || j < 0 || i < 0 {
		return false
	}
	if !visited[i][j] {
		return true
	}
	return false
}
