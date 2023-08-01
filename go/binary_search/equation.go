// 整数 N が与えられます。 次の式を満たす0 以上 100 以下の実数 X を求めてください。
// `X(X(X+1)+2)+3 = N`
// ただし、X は 1 つに定まることがわかっています。
// また、絶対誤差 (想定解と出力の差の絶対値) は 0.01 まで許されます。

package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	float_n := float64(n)

	left := 0.00
	right := 100.00

	for(0.01 < math.Abs(right - left)){
		mid := (right + left) / 2
		equation_answer := mid*(mid*(mid+1)+2)+3

		if(float_n > equation_answer){
			left = mid
		}else{
			right = mid
		}
	}
	fmt.Println(left)
}
