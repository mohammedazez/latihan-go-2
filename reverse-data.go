package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		name, newName string
	)

	name = "muhamad"
	for i := len(name) - 1; i >= 0; i-- {
		//fmt.Println(string(name[i]))
		newName += string(name[i])
	}
	fmt.Println(newName)
	toArray := strings.Fields(newName)
	fmt.Println(toArray)
}
