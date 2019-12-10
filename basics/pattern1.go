<<<<<<< HEAD
package main

import "fmt"

func main() {
	var n int = 0
	fmt.Println("Enter the range")
	fmt.Scan(&n)
	fmt.Println()
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
=======
package main

import "fmt"

func main() {
	var n int = 0
	fmt.Println("Enter the range")
	fmt.Scan(&n)
	fmt.Println()
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
>>>>>>> da444e204ebd8ddc847816e79e4c3bc1e147a7e9
