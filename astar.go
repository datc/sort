package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"time"
)

func main() {
	fmt.Println("ASTAR\t\t")
	display(data)
	dfs(startVst.i, startVst.j)
}

var (
	data = [][]int{
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 7, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0},
	}
	visited   [][]bool
	startV    = 7
	targetV   = 9
	startVst  = vst{i: 1, j: 1}
	obstacle  = 1
	targetVst = vst{i: 6, j: 6}
	m         = len(data)
	n         = len(data[0])
	start     = time.Now()
)

func init() {
	if m <= 0 {
		return
	}
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	data[targetVst.i][targetVst.j] = targetV
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
	if v == obstacle {
		return
	}
	fmt.Println(fmt.Sprintf("(%d,%d):%d\t\t", i, j, v))
	if !startVst.reached(i, j) {
		data[i][j] = 2
	}
	display(data)
	visited[i][j] = true
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
reached the some position
*/
func (v vst) reached(i, j int) bool {
	return i == v.i && j == v.j
}

/**
distance to the target position
*/
func (v vst) dis(trgt vst) int {
	return (int(math.Abs(float64(trgt.i-v.i)) + math.Abs(float64(trgt.j-v.j))))
}

type vstSlc []vst

func (v vstSlc) Len() int {
	return len(v)
}

func (v vstSlc) Less(i, j int) bool {
	return v[i].dis(targetVst) < v[j].dis(targetVst)
}

func (v vstSlc) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v vstSlc) Sort() {
	sort.Sort(v)
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
	vs := vstSlc(vsts)
	vs.Sort()
	_ = vs
	return vsts
}

/**
check the next step if safe: out of range or visited
*/
func safe(i, j int) bool {
	if i >= m || j >= n || j < 0 || i < 0 {
		return false
	}
	if targetVst.reached(i, j) {
		fmt.Println("target is found, costs ", time.Since(start))
		os.Exit(0)
	}
	if !visited[i][j] {
		return true
	}
	return false
}
