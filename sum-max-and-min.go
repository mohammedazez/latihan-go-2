package main

import (
	"fmt"
)

func main() {
	var number = [...]int{10, 120, 14, 50, 5}

	smallesNumber := number[0]
	biggestNumber := number[0]

	for _, element := range number {
		if element < smallesNumber {
			smallesNumber = element
		}

		if element > biggestNumber {
			biggestNumber = element
		}
	}

	fmt.Println("Ini adalah array terkecil", smallesNumber)
	fmt.Println("Ini adalah array terbesar", biggestNumber)

	postiveSum := biggestNumber + smallesNumber
	negativeSum := biggestNumber - smallesNumber
	result := postiveSum + negativeSum

	fmt.Println(postiveSum)
	fmt.Println(negativeSum)
	fmt.Println(result)
}
