package main

import "os"

func main() {
	fileNames := os.Args[1:]
	if fileNames == nil {
		return
	}
	concurrentArchivator(fileNames)
	sequencedArchivator(fileNames)
}