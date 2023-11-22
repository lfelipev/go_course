package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f_name := os.Args[1]

	f, err := os.Open(f_name)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
