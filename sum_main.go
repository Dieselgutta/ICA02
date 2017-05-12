package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ArgsA := os.Args[1]
	ArgsB := os.Args[2]
	i, _ := strconv.ParseFloat(ArgsA, 64)
	j, _ := strconv.ParseFloat(ArgsB, 64)

	fmt.Println(i + j)
}
