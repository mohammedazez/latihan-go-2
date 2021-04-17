package main

import (
	"fmt"
	"strings"
)

func main()  {
	//var (
	//	name, newName string
	//)
	//
	//name = "muhamad"
	//for index, nama := range toArray {
	//	fmt.Println(index, nama)
	//}
	//for i := len(name)- 1; i >= 0; i-- {
	//	//fmt.Println(string(name[i]))
	//	newName += string(name[i])
	//}
	//fmt.Println(newName)
	//
	//toArray := strings.Fields(newName)
	//fmt.Println(toArray)

	// A slice of 3 strings.
	values := []string{"a", "b", "c"}
	// Join with no separator to create a compact string.
	fmt.Println(values)
	joined := strings.Join(values, ",")
	fmt.Println(joined)
	toArray := strings.Fields(joined)
	fmt.Println(toArray)
}