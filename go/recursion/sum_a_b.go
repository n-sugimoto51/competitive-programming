// A から B までの整数の総和
// 1 から 5 までの総和
package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	cast_a := int(a)
	cast_b := int(b)

	sum := calc_sum(cast_a, cast_b)
	fmt.Println(sum)
}

func calc_sum(cast_a int, cast_b int) int {
	if cast_a == cast_b {
		return cast_b
	} else {
		x := cast_b - 1
		return calc_sum(cast_a, x) + cast_b
	}
}
