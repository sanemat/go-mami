package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Version of mami
const Version = "0.1.0"

func mami(current time.Time, givenDir string) (string, error) {
	// memo: basename = time.strftime('%Y-%m-%d-%H-%M-%S')
	const timeFormat = "2006-01-02-15-04-05"

	// ErrDir is the error resulting if no directory given
	ErrDir := errors.New("mami requires MAMI_DIR or directory option")

	baseName := current.Format(timeFormat)
	var targetDir string
	if givenDir != "" {
		targetDir = givenDir
	} else {
		// NOTE: add document about MAMI_DIR
		targetDir = os.Getenv("MAMI_DIR")
	}
	if targetDir == "" {
		return "", ErrDir
	}

	return filepath.Join(targetDir, baseName+".txt"), nil
}

func main() {
	givenDir := flag.String("d", "", "target directory")
	showVersion := flag.Bool("v", false, "show version")

	flag.Parse()
	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}
	current := time.Now()
	if output, err := mami(current, *givenDir); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(output)
	}
}
