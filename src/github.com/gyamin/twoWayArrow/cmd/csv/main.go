package main

import (
	"fmt"
	"os"
)

func main() {
	p, _ := os.Getwd()
	fmt.Println(p)
}
