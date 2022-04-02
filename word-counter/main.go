package main

// bufio to read text
// fmt to print "formatted" output
// io to provide io.Reader interface
// os to use operating system resources
// flag package helps manage command line flags such as "-l"

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main () {
	// this defines a boolean flag -l to count lines instead of words
	// default value is false for this case
	lines := flag.Bool("l", false, "Count lines")
	
	fmt.Println(count(os.Stdin, *lines))
}

func count (r io.Reader, countLines bool) int {
	// this scanner variable "reads" text from a Reader
	scanner := bufio.NewScanner(r)
	// define the scanner split type to "words"
	// if countline is false, this is what we want to do
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	// this defines word count
	wc := 0
	for scanner.Scan() {
		wc++
	}
	return wc
}


