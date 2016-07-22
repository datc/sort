package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ASTAR")
	display(data)
	dfs(1,1)
	//display(data)
}

var (
	data = [][]int{
		[]int{0,0,0,0,0,0,0,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,1,0,9,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,1,0,0,0},
		[]int{0,0,0,0,0,0,0,0},
	}
	visited [][]bool
	target = 9
	m=len(data)
	n=len(data[0])
)

func init() {
	if m <= 0 {
		return
	}
	visited = make([][]bool,m)
	for i := 0; i < m; i++ {
		visited[i]=make([]bool,n)
	}
}

func display(data [][]int)  {
	for _,it:= range data{
		for _,itt:=range it{
			fmt.Print(itt,"  ")
		}
		fmt.Println()
	}
}

func rute(r [][]bool)  {
	for _,it:= range r{
		for _,itt:=range it{
			fmt.Print(itt,"  ")
		}
		fmt.Println()
	}
}

func dfs(i,j int)  {
	v := data[i][j]
	if v==target {
		fmt.Println("target is found.")
		os.Exit(0)
	}
	if visited[i][j] || v!=0 {
		return
	}
	fmt.Println(fmt.Sprintf("(%d,%d):%d",i,j,v))
	display(data)
	visited[i][j]=true
	data[i][j]=2
	dfs(next(i,j))
	visited[i][j]=false
	data[i][j]=0
}

func next(i, j int) (int,int) {
	if safe(i-1,j) {
		return i-1,j
	}
	if safe(i+1,j) {
		return i+1,j
	}
	if safe(i, j - 1) {
		return i,j-1
	}
	if safe(i, j + 1) {
		return i,j+1
	}
	return i,j
}

func safe(i,j int)bool  {
	if i>=m||j>=n|| j<0||i<0 {
		return false
	}
	if !visited[i][j] {
		return true
	}
	return false
}