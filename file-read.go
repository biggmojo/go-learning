package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil" // Deprecated in Go 1.16
	"os"
)

func main() {
	// Read entire file
	f1, err1 := ioutil.ReadFile("read-me.txt")
	if err1 != nil {
		fmt.Println("Error")
	}
	fmt.Println(string(f1))

	// Read a chunk of a file
	f2, err2 := os.Open("read-me.txt")
	if err2 != nil {
		fmt.Println("Error  opening the file.")
	}

	const chunkSize = 4 // Chunk size

	buf := make([]byte, chunkSize) // Buffer

	for {
		// Read content to buffer
		readTotal, err3 := f2.Read(buf)
		if err3 != nil {
			if err3 != io.EOF {
				fmt.Println(err3)
			}
			break
		}
		fmt.Println(string(buf[:readTotal])) // Print content from buffer
	}

	// Read a file line by line
	f4, err4 := os.Open("read-me.txt")
	if err4 != nil {
		fmt.Println(err4)
	}

	scanner := bufio.NewScanner(f4)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err4 := scanner.Err(); err4 != nil {
		fmt.Println(err4)
	}

	f4.Close()

	// Read file word by word
	f5, err5 := os.Open("read-me.txt")
	if err5 != nil {
		fmt.Println(err5)
	}

	scanner2 := bufio.NewScanner(f5)
	scanner2.Split(bufio.ScanWords)
	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}

	if err5 := scanner2.Err(); err5 != nil {
		fmt.Println(err5)
	}

	f5.Close()
}
