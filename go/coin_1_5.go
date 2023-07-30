package main

import (
	"fmt"
)

// 1 円玉 と 5 円玉を使って N 円を支払うとき、最低何枚の硬貨が必要ですか。
func main() {
	var n int
	fmt.Scanf("%d", &n)

	coin := 0
	for n >= 5 {
		n -= 5
		coin++
	}
	for n >= 1 {
		n--
		coin++
	}
	fmt.Println(coin)
}


// 模範解答
func answer(){
	var n int
	fmt.Scanf("%d", &n)

	coin_sum := (n / 5) + (n % 5)
	fmt.Println(coin_sum)
}
