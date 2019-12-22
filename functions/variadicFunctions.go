//program to find whether an integer exists in an input list
package main

import (
	"fmt"
)

func main() {
	find(10,10,20,30);
	find(1,45,3,74,83,44,2);
	find(43,64,83,43);
	find(45);
}

func find(num int, list ... int) {
	fmt.Printf("type of list is %T\n", list);
	found := false;
	for i, v := range list {
		if v == num {
			fmt.Println(num, "found at index", i, "in", list);
			found = true;
		}
	}

	if !found {
		fmt.Println(num, "not found in", list);
	}
	fmt.Printf("\n");
}
