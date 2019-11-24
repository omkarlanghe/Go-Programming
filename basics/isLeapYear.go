package main

import "fmt"

func main() {
	var year int;
	fmt.Println("\nEnter year to check");
	fmt.Scan(&year);

	isLeap := IsLeapYear(year);
	if(isLeap) {
		fmt.Println("Its a Leap Year");
	} else {
		fmt.Println("Its not a Leap Year");
	}
}

func  IsLeapYear(year int) bool {
	if (year % 400) == 0 {
		return true;
	} else if (year % 4) == 0 {
		return true;
	} else if (year % 100) == 0 {
		return false;
	} else {
		return false;
	}
}
