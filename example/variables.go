package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	fmt.Println(s)
	const n = 6000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
	fmt.Println("math.Sin=" + fmt.Sprintf("%f", math.Sin(n)))
	fmt.Println("math.Cos=" + fmt.Sprintf("%f", math.Cos(n)))
	fmt.Println("math.Tan=" + fmt.Sprintf("%f", math.Tan(n)))
}
