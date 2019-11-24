package main

import "fmt"

func main() {
	var num int;
	fmt.Println("\nEnter the number for factorial");
	fmt.Scan(&num);

	fact := 1;
	for i := 2 ; i <= num ; i++ {
		fact = fact * i;
		fmt.Println("\nFactorial : ",fact);
	}
}
