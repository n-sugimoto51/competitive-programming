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
	
}
