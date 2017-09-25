package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	FILE_PATH    = "/data/ShareFile.txt"
	FILESIZE     = 0
	EVERYMINUTE  = 2
	DEFSLEEPTIME = 2
)

func main() {
	fStat, err := os.Stat(FILE_PATH)
	if err != nil {
		panic(err.Error())
	}

	fHandle, err := os.Open(FILE_PATH)
	if err != nil {
		panic(err.Error())
	}

	for {
		if fStat.Size() <= int64(FILESIZE) {
			fmt.Println("Null Subscription ID is found!!!!")
			time.Sleep(DEFSLEEPTIME * time.Minute)
			continue
		}

		if fHandle == nil {
			fHandle, err = os.Open(FILE_PATH)
			if err != nil {
				panic(err.Error())
			}
		}
		scanner := bufio.NewScanner(fHandle)

		// Read subscriptionID per line, For 10 subscription call a go routine
		for scanner.Scan() {
			fmt.Println("In Client = ", scanner.Text())
		}

		//close file
		fHandle.Close()
		fHandle = nil

		time.Sleep(EVERYMINUTE * time.Minute)

	}
}
