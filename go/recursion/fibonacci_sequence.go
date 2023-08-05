// フィボナッチ数列を求める
// 入力 n

// 以下のように再帰関数を使ってフィボナッチ数列(1 から N までの整数の総和)を求めています
// 以下のコードを中学生が理解できるように説明してください

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
