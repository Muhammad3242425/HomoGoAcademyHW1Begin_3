package main

import (
	"fmt"
)

func main() {
	var a int
	var b int

	fmt.Print("a= ")
	fmt.Scan(&a)
	fmt.Print("b= ")
	fmt.Scan(&b)

	fmt.Println("S= ", a*b)
	fmt.Println("P= ", 2*(a+b))

}
