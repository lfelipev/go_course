package main

import (
	"fmt"
	"io"
	"os"
)

//go build main.go
//main.exe main.go

func main() {
	f_name := os.Args[1]

	f, err := os.Open(f_name)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
