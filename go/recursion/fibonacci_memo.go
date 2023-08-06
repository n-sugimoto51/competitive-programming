// フィボナッチ数列を求める
// 入力 n

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	// momoスライスの要素全てを-1で初期化
	memo := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = -1
	}


	fibonacci := calc_fibonacci_memo(n, &memo)
	fmt.Println(fibonacci)
}

func calc_fibonacci_memo(n int, memo *[]int) int {

	// メモにデータがあればそれを返す
	if (*memo)[n] != -1 {
		return (*memo)[n]
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		x := n - 1
		y := n - 2
		calc_result := calc_fibonacci_memo(x, memo) + calc_fibonacci_memo(y, memo)
		(*memo)[n] = calc_result
		return calc_result
	}
}
