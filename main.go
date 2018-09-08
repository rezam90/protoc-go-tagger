package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	var inputFile string
	var xxxTags string
	flag.StringVar(&inputFile, "input", "", "path to input file")
	flag.StringVar(&xxxTags, "skip_xxx", "", "skip tags to inject on XXX fields")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	matches, err := filepath.Glob(inputFile)
	if err != nil {
		log.Fatal("can't read files", err.Error())
	}

	for _, filePath := range matches {
		areas, err := parseFile(filePath, xxxSkipSlice)
		if err != nil {
			log.Fatal(err)
		}
		if err = writeFile(filePath, areas); err != nil {
			log.Fatal(err)
		}
	}
}
