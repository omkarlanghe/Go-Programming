package main

import "fmt"

func main() {
	var n int = 0
	fmt.Println("Enter the range")
	fmt.Scan(&n)
	fmt.Println()
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
