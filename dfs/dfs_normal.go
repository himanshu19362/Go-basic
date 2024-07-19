package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func dfs(node int, adjList *[][]int) {
	for _, x := range (*adjList)[node] {
		dfs(x, adjList)
	}
}

func main() {
	adjList := make([][]int, 1690)
	for i := range adjList {
		adjList[i] = make([]int, 0)
	}
	createAdjList(&adjList)
	start := time.Now()
	dfs(1, &adjList)
	end := time.Now()

	timeDiff := end.Sub(start)
	fmt.Println("%v\n", timeDiff)
}

func createAdjList(adjList *[][]int) {
	file, _ := os.Open("graph.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineNo := 1
	for scanner.Scan() {
		for _, element := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.Atoi(element)
			(*adjList)[lineNo] = append((*adjList)[lineNo], num)
		}
		lineNo++
	}
}
