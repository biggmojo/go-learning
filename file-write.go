package main

import "fmt"
import "os"

func main() {
	// Create the file
	f, err := os.Create("writeme.txt")
	if err != nil {
		fmt.Println(err)
	}

	// Write directly into the file
	f.Write([]byte("Whale hello there."))

	// Write a string
	f.WriteString("\nThis is a written string.")

	f.Close()

	// Append to an existing file
	f2, err2 := os.OpenFile("writeme.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	newLine := "\nThis was appended to the file after creation."
	_, err2 = fmt.Fprintln(f2, newLine)
	if err2 != nil {
		fmt.Println(err2)
	}
}
