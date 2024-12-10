package main

import "fmt"

func main() {
	var billAmount int
	fmt.Print("Enter the total bill amount: ")
	fmt.Scan(&billAmount)

	fmt.Println(billAmount)
}
