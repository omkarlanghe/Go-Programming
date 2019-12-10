package main

import "fmt"

func main() {
	var arr [10]int;
	fmt.Println("Enter the elements in an array");
	for i := 0 ; i < 10 ; i++ {
		fmt.Scan(&arr[i]);
	}

	digitSum, elementSum := SumOperationOnArray(arr);

	fmt.Println("Sum of digits from total elements is : ", digitSum);
	fmt.Println("Sum of elements from array is : ", elementSum);
}

func SumOperationOnArray(arr[10] int) (int, int) {
	var temp, rem, sum, arraySum int;

	for i := 0 ; i < len(arr) ; i++ {
		temp = arr[i];

		for ; temp != 0 ; {
			rem = temp%10;
			sum += rem;
			temp = temp/10;
		}
		arraySum += arr[i];
	}
	return sum, arraySum;
}
