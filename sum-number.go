package main

import "fmt"

func main(){
	var (
		number = [...]int{1, 2, 3, 4, 5}
		newNumber int
	)
	for i := 0; i <= len(number); i++ {
		fmt.Println("ini number", number)
	newNumber += i
		fmt.Println("ini angka", newNumber)
	}
}


