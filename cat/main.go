package main

import (
	"io"
	"log"
	"os"
)

func printFile(f *os.File) {
	buf := make([]byte, 4096)
	for {
		read, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalf("Unexpected error during read a file: file = %s, error = %v", f.Name(), err)
			}
		} else {
			os.Stdout.Write(buf[:read])
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Arguments amount should be 2")
	}
	fName := os.Args[1]
	f, err := os.Open(fName)
	if err != nil {
		log.Fatalf("Unexpected error during open a file: name = %s, error = %v", fName, err)
	}
	defer f.Close()
	printFile(f)
}
