package main

import (
	"fmt"
	"sync"

	"example.com/hunaidsav/FileTypeAnalyzer/algorithms"
	"example.com/hunaidsav/FileTypeAnalyzer/ds"
)

func doWork(path string, patterns ds.ByPriority) {
	text := string(readFile(path))
	matched := false
	for _, p := range patterns {
		if algorithms.ContainsKMP(text, p.Pattern) {
			matched = true
			fmt.Println(path, ":", p.DocumentType)
			break
		}
	}
	if !matched {
		fmt.Println(path, ": Unknown File Type")
	}
}

func Analyzer(filesDir, patternsDir string) {

	patterns := getPatterns(patternsDir)

	filePaths := getAllFileNames(filesDir)
	n := len(filePaths)

	var wg sync.WaitGroup
	wg.Add(n)
	for _, path := range filePaths {

		go func(p string) {
			doWork(p, patterns)
			wg.Done()
		}(path)
	}

	wg.Wait()
}

func main() {

	Analyzer("/home/brian/Documents/", "/home/brian/go/src/hello/patterns.db")
}
