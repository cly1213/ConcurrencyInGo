package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type PersonName struct {
	firstName string
	lastName  string
}

func main() {
	tmpFileName := ""
	fmt.Println("Enter File Name :")
	fmt.Scanln(&tmpFileName)
	file, err := os.Open(tmpFileName)

	if err != nil {
		fmt.Println("[ERROR] unable to read file: ", err)
	}

	defer file.Close()

	var personNames []PersonName
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = string(scanner.Text())
		var tfname = strings.Split(line, " ")[0]
		var tlname = strings.Split(line, " ")[1]

		personName := PersonName{
			firstName: tfname,
			lastName:  tlname,
		}

		personNames = append(personNames, personName)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, personName := range personNames {
		fmt.Println("----------------------------")

		fmt.Println("First Name:", personName.firstName)
		fmt.Println("Last  Name:", personName.lastName)

	}

}
