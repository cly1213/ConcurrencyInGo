package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	p := make(map[string]string)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Input a name:")

	name, _ := reader.ReadString('\n')
	mname := string(name)
	p["name"] = mname[0 : len(mname)-1]

	fmt.Printf("Input an address:")
	address, _ := reader.ReadString('\n')
	maddress := string(address)
	p["address"] = maddress[0 : len(maddress)-1]

	data, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		fmt.Println("Error:", err)
	}

	//os.Stdout.Write(b)
	fmt.Printf("JSON Data : %v", string(data))
}
