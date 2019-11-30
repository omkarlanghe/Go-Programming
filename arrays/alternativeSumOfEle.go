package main

import "fmt"

func main() {
	var arr [10]int;

	fmt.Println("Enter the elements in an array");
	for i := 0 ; i < 10 ; i++ {
		fmt.Scan(&arr[i]);
	}

	difference := FindDifference(arr);

	fmt.Println("Difference is : ", difference);
	/*fmt.Println("Elements in array");
	for _, v := range arr {
		fmt.Println("element -> ", v);
	}*/
}

func FindDifference(arr[10] int) int {
	var firstsum, secondsum int;
	for i := 0 ; i < len(arr) ; i = i+2 {
		firstsum += arr[i];
		secondsum += arr[i+1];
	}

	if firstsum > secondsum {
		return firstsum - secondsum;
	} else {
		return secondsum - firstsum;
	}

}
