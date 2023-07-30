// 50 円玉, 10円玉, 5 円玉, 1円玉がそれぞれ A0,A1,A2,A4枚あります。
// これらを用いて X 円を支払うとき、最小で何枚の硬貨が必要ですか。 ただし、支払い方は少なくとも 1 つ存在することが保証されています。
// 入力
// 22
// 3 1 4 2
package main

import (
	"fmt"
)

func main() {
	var x, a, b, c, d int
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	coin_count := 0

	coin_calc(&x, &a, &coin_count, 50)
	coin_calc(&x, &b, &coin_count, 10)
	coin_calc(&x, &c, &coin_count, 5)
	coin_calc(&x, &d, &coin_count, 1)

	fmt.Println(coin_count)
}

func coin_calc(x *int, n *int, coin_count *int, coin_num int) {
	if(*x == 0){
		return
	}
	max_n := *x / coin_num
	if(max_n > 0 && *n > 0){
		min_n := 0
		if(max_n > *n){
			min_n = *n
		}else{
			min_n = max_n
		}
		*x -= min_n * coin_num
		*coin_count += min_n
	}
}


