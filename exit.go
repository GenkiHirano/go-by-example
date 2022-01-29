package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println()

	defer fmt.Println("!")

	os.Exit(3)
}
