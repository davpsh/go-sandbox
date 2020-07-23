package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var rootPath string
var enableDebugMessage bool

func main() {
	flag.StringVar(&rootPath, "path", "./test_data", "path to folder")
	flag.BoolVar(&enableDebugMessage, "debug", false, "enable debug messages or not")
	flag.Parse()

	totalSum := findTotalSum(rootPath)
	fmt.Println("Total sum is ", totalSum)
}

func findTotalSum(path string) int64 {
	var totalSum int64 = 0
	paths := findCountFiles(path)

	for _, path := range paths {
		totalSum += countTotalSum(path)
	}

	return totalSum
}

func findCountFiles(path string) []string {
	paths := make([]string, 0)

	filepath.Walk(path, func(walkPath string, info os.FileInfo, err error) error {
		if walkPath == path {
			return nil
		}

		if !info.IsDir() && info.Name() == "count" {
			paths = append(paths, walkPath)
		}

		return nil
	})

	return paths
}

func countTotalSum(path string) int64 {
	var totalSum int64 = 0

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Printf("open file error: %v\n", err)
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineValue, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			lineValue = 0
		}
		totalSum += lineValue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if enableDebugMessage {	fmt.Printf("Path: %v\n", path)
		fmt.Printf("Total sum: %d\n", totalSum)
		fmt.Printf("=================\n")

	}

	return totalSum
}
