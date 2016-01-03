package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

// memo: basename = time.strftime('%Y-%m-%d-%H-%M-%S')
const timeFormat = "2006-01-02-15-04-05"

var givenDir = flag.String("d", "", "target directory")

func mami(current time.Time, givenDir string) (string, error) {
	baseName := current.Format(timeFormat)
	return givenDir + baseName + ".txt", nil
}

func main() {
	flag.Parse()
	current := time.Now()
	if output, err := mami(current, *givenDir); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
