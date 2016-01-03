package main

import (
	"fmt"
	"log"
	"time"
)

// basename = time.strftime('%Y-%m-%d-%H-%M-%S')
const timeFormat = "2006-01-02-15-04-05"

func mami(current time.Time) (string, error) {
	baseName := current.Format(timeFormat)
	return "foo/" + baseName + ".txt", nil
}

func main() {
	current := time.Now()
	if output, err := mami(current); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
