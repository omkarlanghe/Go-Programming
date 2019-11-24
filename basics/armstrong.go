package main

import "fmt"

func main() {
	var num int = 0;
	fmt.Println("\nEnter the number to check");
	fmt.Scan(&num);

	retVal := ArmstrongNum(num);
	if retVal == num {
		fmt.Println("Its an armstrong number");
	} else {
		fmt.Println("Not an armstrong number");
	}
}

func ArmstrongNum(num int) int {
	var rem, result int = 0, 0;
	for ;num != 0; {
		rem = num % 10;
		result = result + (rem*rem*rem);
		num = num / 10;
	}
	return result
}
