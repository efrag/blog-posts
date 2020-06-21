package main

import (
	"flag"
	"fmt"
	"os"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "\nStructor requires the name of a directory as an input\n")
	fmt.Fprintf(os.Stderr, "\tstructor -domain D [single directory name]\n")
	fmt.Fprintf(os.Stderr, "\n")
}

// basic setup for the structor command
func setup() string {
	dir := flag.String("domain", "", "A domain package directory")
	flag.Usage = Usage
	flag.Parse()

	if *dir == "" {
		flag.Usage()
		os.Exit(2)
	}

	return *dir
}

func main() {
	dir := setup()

	st := NewStructor(dir)
	err := st.Generate()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to generate accessor files: %v \n", err)
		os.Exit(2)
	}
}
