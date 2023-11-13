// lets start
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scanf("%f\n%f", &a, &b)
	fmt.Printf("%f", math.Hypot(a,b))
}
