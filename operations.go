package main

import (
	"archive/zip"
	"io"
	"os"
	"sync"
)

func sequencedArchivator(fileNames []string) {
	for _, fileName := range fileNames {
		fileForArchive(fileName)
	}
}

func concurrentArchivator(fileNames []string) {
	waitGroup := sync.WaitGroup{}
	for _, val := range fileNames {
		waitGroup.Add(1)
		go func(wg *sync.WaitGroup, fileName string) {
			defer func() {
				waitGroup.Done()
			}()
			fileForArchive(val)
		}(&waitGroup, val)
	}
	waitGroup.Wait()
}




func fileForArchive(fileName string) {
	zipFile, err := os.Create(archivesPath + fileName + zipFormat)
	if err != nil {
		return
	}

	defer func() {
		err = zipFile.Close()
		if err != nil {
			return
		}
	}()

	writer := zip.NewWriter(zipFile)

	defer func() {
		err = writer.Close()
		if err != nil {
			return
		}
	}()

	file, err := os.Open(filesPath + fileName)
	if err != nil {
		return
	}

	archive, err := writer.Create(fileName)
	if err != nil {
		return
	}
	_, err = io.Copy(archive, file)
	if err != nil {
		return
	}
}