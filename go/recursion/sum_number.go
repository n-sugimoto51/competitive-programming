// 1 から N までの整数の総和
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	cast_n := int(n)

	sum := calc_sum(cast_n)
	fmt.Println(sum)
}

func calc_sum(cast_n int) int {
	if cast_n == 0 {
		return 0
	} else {
		x := cast_n - 1
		return calc_sum(x) + cast_n
	}
}
