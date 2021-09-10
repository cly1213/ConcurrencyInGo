package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter your string: ")

	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}

	line = strings.TrimSpace(line)
	lowerString := strings.ToLower(line)
	formattedString := strings.ReplaceAll(lowerString, " ", "")

	if strings.Contains(formattedString, "a") { // search 'a'
		if strings.HasPrefix(formattedString, "i") { // search 'i' a start
			if strings.HasSuffix(formattedString, "n") { // search 'n' at end
				fmt.Println("Found!")
				os.Exit(0) // exits code if found.
			}
		}
	}
	fmt.Println("Not Found!")
}
