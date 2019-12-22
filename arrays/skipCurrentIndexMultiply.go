package main

import "fmt"

func main() {
	var arr[5]int;

	fmt.Println("Enter the elements in an array");
	for i := 0 ; i < len(arr) ; i++ {
		fmt.Scan(&arr[i]);
	}

	skipIndexAndMultiply(arr);
}

func skipIndexAndMultiply(arr [5]int) {
	var res[5]int;
	for i := 0 ; i < len(arr) ; i++ {
		temp := 1;
		for j := 0 ; j < len(arr) ; j++ {
			if i == j {
				continue;
			}
			temp = temp * arr[j];
			res[i] = temp;
		}
	}

	for i := 0 ; i < len(res) ; i++ {
		fmt.Println("result : ", res[i]);
	}
}
