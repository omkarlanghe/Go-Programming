package main

import "fmt"

func main() {
	var fNum, sNum, tNum int;
	fmt.Println("\nEnter 3 numbers");
	fmt.Scan(&fNum);
	fmt.Scan(&sNum);
	fmt.Scan(&tNum);

	maxNum := MaxNum(fNum,sNum,tNum);
	fmt.Println("\nMaximum number is : ",maxNum);
}

func MaxNum(fNum, sNum, tNum int) int {
	if fNum > sNum && fNum >tNum {
		return fNum;
	} else if sNum > tNum {
		return sNum;
	} else {
		return tNum;
	}
}
