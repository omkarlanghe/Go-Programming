package main

import "fmt"

func main() {
	var num int;
	fmt.Println("\nEnter the number");
	fmt.Scan(&num);

	retVal := Pallindrome(num);
	if retVal == num {
		fmt.Println("\nYes its a Pallindrome");
	} else {
		fmt.Println("\nNops its not a Pallindrome");
	}
}

func Pallindrome(num int) int {
	var rem, result int;
	for ;num > 0; {
		rem = num % 10;
		result = result * 10 + rem;
		num = num / 10;
	}
	return result;
}
