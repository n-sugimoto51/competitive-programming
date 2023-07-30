// アルルは N 個のお菓子を持っており、毎日次のいずれかのプランでお菓子を食べます。
// プラン 1：お菓子を 1 つ食べる。
// プラン 2：半分の個数のお菓子を食べる。
// 全てのお菓子を食べ切るまで最短で何日かかるかを答えてください。
// ただし、プラン 2 はお菓子の個数が偶数の場合のみ行うことができます。

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	day_count := 0
	for n > 0 {
		if(n % 2 == 0){
			n = n / 2
		}else{
			n--
		}
		day_count++
	}
	fmt.Println(day_count)
}
