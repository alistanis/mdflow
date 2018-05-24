package main

import (
	"fmt"
	"os"

	"github.com/alistanis/mdflow"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	return mdflow.ExecuteRoot()
}
