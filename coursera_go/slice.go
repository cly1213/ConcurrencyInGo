package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	ListSize int    = 3
	ExitCode string = "X"
)

func main() {

	var input string
	var initSlice = make([]int, 0, ListSize)

	for {
		fmt.Print("Enter an integer", ": ")
		_, err := fmt.Scan(&input)
		exitCode := strings.ToUpper(input)

		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		if exitCode == ExitCode {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		parseIntScanValue, err := strconv.Atoi(input)
		initSlice = append(initSlice, parseIntScanValue)
		sort.Ints(initSlice)
		fmt.Println(initSlice)
	}
}
