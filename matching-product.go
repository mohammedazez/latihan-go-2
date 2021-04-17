package main

import "fmt"

//belum selesai
func main() {
	// 10, 10, 20, 20, 10, 30, 50, 10, 20
	// cari integer yang sama = 10, 10, 20, 20, 10, 10, 20
	// pasangkan 2 yang sama = 10 and 10, 20 and 20, 10 and 10
	// length yang mempunyai pasangan

	product := []int{10, 10, 20, 20, 10, 30, 50, 10, 20}

	sameProduct := []int{}

	for i := range product {
		if product[i] == product[i] {
			sameProduct = append(sameProduct, product[i])
		}
	}

	fmt.Println(sameProduct)

	//for i := 0 ; i < len(product) ; i++ {
	//	//fmt.Println(product[i])
	//	for j := 0 ; j < len(product) ; j++ {
	//		//fmt.Println(product[j])
	//
	//		fmt.Println(product[i] == product[j])
	//
	//		//if product[i] == product[j] {
	//		//	fmt.Println(product[i] )
	//		//}
	//	}
	//}
}
