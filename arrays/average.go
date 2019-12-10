package main

import "fmt"

func main() {
	var arr [10]int;

	fmt.Println("Enter the elements in an array");
	for i := 0 ; i < 10 ; i++ {
		fmt.Scan(&arr[i])
	}

	average := AverageOfEle(arr);
	fmt.Println("Average is : ", average);

}

func AverageOfEle(arr [10]int) int {
	var sum int;
	total := len(arr);
	for i := 0 ; i < len(arr) ; i++ {
		sum += arr[i];
	}

	return(sum/total);
}


