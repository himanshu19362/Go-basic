package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	generateFiles("file1.txt", 100)
	generateFiles("file2.txt", 100)
	generateFiles("file3.txt", 100)
	var wg sync.WaitGroup
	wg.Add(3)
	start := time.Now()
	finalArray := []int{}
	go readFile("file1.txt", &finalArray, &wg)
	go readFile("file2.txt", &finalArray, &wg)
	go readFile("file3.txt", &finalArray, &wg)

	wg.Wait()
	sort.Slice(finalArray, func(i, j int) bool {
		return finalArray[i] < finalArray[j]
	})
	writeToFile("file4.txt", &finalArray)
	end := time.Now()
	timeDiff := end.Sub(start)
	fmt.Println(timeDiff)
}

func generateFiles(filename string, numCount int) {
	file, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	for i := 0; i < numCount; i++ {
		_, err := fmt.Fprintln(file, rand.Intn(1000))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("File created successfully")
}

func readFile(filename string, finalArray *[]int, wg *sync.WaitGroup) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	defer wg.Done()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		*finalArray = append(*finalArray, num)
	}
}

func writeToFile(filename string, finalArray *[]int) {
	file, _ := os.Create(filename)
	defer file.Close()
	for i := 0; i < len(*finalArray); i++ {
		num := (*finalArray)[i]
		fmt.Fprintln(file, num)
	}
}
