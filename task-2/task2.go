package main

import "fmt"

func main() {
	var a, b, c int64
	fmt.Scanf("%v", &a)
	fmt.Scanf("%v", &b)
	fmt.Scanf("%v", &c)

	if (a+b > c) && ((a + c) > b) && ((b + c) > a) {
		fmt.Printf("YES")
	} else {
		fmt.Printf("NO")
	}
}
