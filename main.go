package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b int64
	fmt.Println("Введите диапазон")
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := Array(a, b)
	fmt.Println(c)
}

func Array(a, b int64) (arr []int64) {

	for i := a; i <= b; i++ {
		if big.NewInt(a).ProbablyPrime(0) {
			arr = append(arr, a)
		}
	}
	return
}
