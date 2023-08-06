// フィボナッチ数列を求める
// 入力 n

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	fibonacci := calc_fibonacci(n)
	fmt.Println(fibonacci)
}

func calc_fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		x := n - 1
		y := n - 2
		return calc_fibonacci(x) + calc_fibonacci(y)
	}
}
