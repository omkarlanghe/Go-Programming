package main

import "fmt"

func main() {
	var num1 int = 0;
	var num2 int = 0;

	fmt.Println("Enter first number");
	fmt.Scan(&num1);
	fmt.Println("Enter second number");
	fmt.Scan(&num2);

	result := SumOfTwoIntegers(num1,num2); //short hand assignment
	sub := SubtractionOfTwoInts(num1,num2);
	sumResult, subResult := ReturnMultipleValues(num1,num2);
	sumRes, subRes := ReturningNamedValues(num1,num2);

	addition, _ := ReturnMultipleValues(num1,num2); //I only want sum from this function call, hence used blank identifier '_'

	fmt.Println("\nSum of two integers is: ", result);
	fmt.Println("\nSub of two integers is: ", sub);
	fmt.Println("\nSum is: , Sub is: ",sumResult, subResult);
	fmt.Println("\nNamed Sum: ,Named Sub: ",sumRes, subRes);
	fmt.Println("\nOnly Addition: ",addition);
}

func SumOfTwoIntegers(num1 int, num2 int) int {
	return(num1+num2);
}

func SubtractionOfTwoInts(num1, num2 int) int { //if parameters of same type, just write type once..
	if(num1 > num2) {
		return(num1 - num2);
	} else {
		return(num2 - num1);
	}
}

func ReturnMultipleValues(num1, num2 int) (int, int) { //This is how we return multiple values from single function
	var sub int;
	sum := num1 + num2;
	if(num1 > num2) {
		sub = num1 - num2;
	} else {
		sub = num2 - num1;
	}
	return sum,sub;
}

func ReturningNamedValues(num1, num2 int) (sum, sub int) {
	/*sum and sub are named values and are declared during function declaration, 
	as soon as function body encounters return statement, 
	sum and sub variables are automatically returned */
	sum = num1 + num2;
	sub = num1 - num2;
	return;
}


