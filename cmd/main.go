package main

import (
	"fmt"
)

func ConstantBigO() {
	arr := [4]int{1, 2, 3, 4}
	fmt.Println(arr[1])

	num := arr[len(arr)-1]
	fmt.Println("num", num)

	arr[0] = arr[0] + 10
	fmt.Println(arr[0])
}

func main() {
	fmt.Println("Welcome - writing in terminal")
	ConstantBigO()
}
