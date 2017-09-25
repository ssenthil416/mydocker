package main

import (
	"fmt"
	"os"
	"time"
)

const (
	FILE_PATH = "/data/ShareFile.txt"
)

func main() {
	min := time.Now().Minute() - 1
	for { // for ever
		//sleep till one min
		if min == time.Now().Minute() {
			time.Sleep(1 * time.Minute)
		}
		min = time.Now().Minute()
		subsIds := "SUNDAR--" + time.Now().String()

		//Write subsID to a file
		dataFile, err := os.OpenFile(FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}

		fmt.Println("In Server =", subsIds)
		fmt.Fprintf(dataFile, "%s\n", subsIds)

		dataFile.Close()
	} // for ever}
}
