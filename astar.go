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
	//dfs(startVst.i, startVst.j)
	astar()
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
	visitedV  = 2
	pathV     = 3
	startVst  = vst{i: 1, j: 1}
	obstacle  = 1
	targetVst = vst{i: 4, j: 6}
	m         = len(data)
	n         = len(data[0])
	open      sortedMap
	closeM    map[string]*vst
	start     = time.Now()
	step      = 0
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
	open = sortedMap{vstMap: make(map[vst]*vst)}
	open.add(&startVst)
	closeM = make(map[string]*vst)
}

type sortedMap struct {
	vstMap map[vst]*vst
}

func (m *sortedMap) add(v *vst) {
	//fmt.Println(v,v.parent)
	if v.parent != nil {
		v.step = v.parent.step + 1
	} else {
		v.step = 1
	}
	m.vstMap[*v] = v
}

func (m *sortedMap) remove(v *vst) {
	closeM[v.String()] = v
	delete(m.vstMap, *v)
}

/**
should use bubble sort to find the min(fn)
*/
func (m *sortedMap) Sort() *vst {

	min := 10000
	var minVst *vst
	fn := 0
	for _, it := range m.vstMap {
		fn = it.fn()
		if min > fn {
			min = fn
			minVst = it
		}
	}
	return minVst
}

/**
display map
*/
func display(data [][]int) {
	println()
	for _, it := range data {
		for _, itt := range it {
			fmt.Print(itt, "  ")
		}
		fmt.Println()
	}
}

func astar() {
	i := 0
	for {
		i++
		if i > 65 {
			break
		}
		currentVst := open.Sort()
		if nil == currentVst {
			fmt.Println("no path.")
			return
		}
		fmt.Println(currentVst.hashString(), currentVst.step, "\t\t")
		//fmt.Println(closeM)
		//fmt.Println("open:",open.vstMap)
		//fmt.Println(open.vstMap[*currentVst],closeM[currentVst.String()])
		if currentVst.reached(targetVst.i, targetVst.j) {
			fmt.Println("target is found.")
			return
		}
		open.remove(currentVst)
		if !startVst.reached(currentVst.i, currentVst.j) {
			data[currentVst.i][currentVst.j] = visitedV
		}
		//display(data)
		vsts := currentVst.next()
		for _, it := range vsts {
			if it.reached(targetVst.i, targetVst.j) {
				it.parent = currentVst
				currentVst.path(0)
				display(data)
				os.Exit(0)
			}
			if closeM[it.String()] != nil {
				continue
			}
			if nil == open.vstMap[*it] {
				if it.parent == nil {
					it.parent = currentVst
				}
				open.add(it)
			} else {
				if it.fn() >= it.fn() {
					it.parent = currentVst
				}
			}
		}
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
		data[i][j] = visitedV
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
	i, j   int  // pos
	step   int  // step count
	parent *vst // parent vst
	f      int
}

func (v vst) path(depth int) {
	prnt := v.parent
	if prnt != nil && depth < 65 {
		data[v.i][v.j] = pathV
		prnt.path(depth + 1)
		fmt.Print("-->", v.hashString())
	} else {
		print("start")
	}
}

func (v vst) newOne() *vst {
	return &vst{
		i:      v.i,
		j:      v.j,
		parent: v.parent,
		step:   v.step,
		f:      v.f,
	}
}

func (v vst) fn() int {
	return v.dis(targetVst) + step*8
}

func (v *vst) String() string {
	return fmt.Sprintf("v%d%d", v.i, v.j)
}

/**
reached the some position
*/
func (v vst) reached(i, j int) bool {
	return i == v.i && j == v.j
}

func (v vst) hashString() string {
	return fmt.Sprintf("%d_%d", v.i, v.j)
}

/**
distance to the target position
*/
func (v vst) dis(trgt vst) int {
	return (int(math.Abs(float64(trgt.i-v.i)) + math.Abs(float64(trgt.j-v.j))))
}

func (v vst) next() []*vst {
	i, j := v.i, v.j
	vsts := make([]*vst, 0, 4)
	if safe(i-1, j) {
		vsts = append(vsts, &vst{i: i - 1, j: j})
	}
	if safe(i+1, j) {
		vsts = append(vsts, &vst{i: i + 1, j: j})
	}
	if safe(i, j-1) {
		vsts = append(vsts, &vst{i: i, j: j - 1})
	}
	if safe(i, j+1) {
		vsts = append(vsts, &vst{i: i, j: j + 1})
	}
	return vsts
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

type vstSlc2 []*vst

func (v vstSlc2) Len() int {
	return len(v)
}

func (v vstSlc2) Less(i, j int) bool {
	return v[i].dis(targetVst)+v[i].step < v[j].dis(targetVst)+v[j].step
}

func (v vstSlc2) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v vstSlc2) Sort() {
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
		return true
	}
	if data[i][j] == obstacle {
		return false
	}
	if !visited[i][j] {
		return true
	}
	return false
}
