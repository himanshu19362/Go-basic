package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"
// )

// func dfs(node int, adjList *[][]int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for _, x := range (*adjList)[node] {
// 		wg.Add(1)
// 		go dfs(x, adjList, wg)
// 	}
// }

// func main() {
// 	adjList := make([][]int, 1690)
// 	for i := range adjList {
// 		adjList[i] = make([]int, 0)
// 	}
// 	createAdjList(&adjList)
// 	var wg sync.WaitGroup
// 	start := time.Now()
// 	wg.Add(1)
// 	go dfs(1, &adjList, &wg)
// 	wg.Wait()
// 	end := time.Now()

// 	timeDiff := end.Sub(start)
// 	fmt.Println("%v\n", timeDiff)
// }

// func createAdjList(adjList *[][]int) {
// 	file, _ := os.Open("graph.txt")
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)

// 	lineNo := 1
// 	for scanner.Scan() {
// 		for _, element := range strings.Split(scanner.Text(), " ") {
// 			num, _ := strconv.Atoi(element)
// 			(*adjList)[lineNo] = append((*adjList)[lineNo], num)
// 		}
// 		lineNo++
// 	}
// }
