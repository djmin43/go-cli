package main

// bufio to read text
// fmt to print "formatted" output
// io to provide io.Reader interface
// os to use operating system resources
import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {
	fmt.Println(count(os.Stdin))
}

func count (r io.Reader) int {
	// this scanner variable "reads" text from a Reader
	scanner := bufio.NewScanner(r)

	// define the scanner split type to "words"
	scanner.Split(bufio.ScanWords)

	// this defines word count
	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}


