// N+6 以上 100000 以下
// saving = X * N
// X * saving + 1

package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	cast_n := float64(n)
	cast_m := float64(m)
	left := 0.00
	right := 10.00
	for(0.001 < right - left){
		mid := left + (right - left) / 2 // オーバーフロー回避
		answer := calc_equation(cast_n, mid)
		if(cast_m > answer){
			left = mid
		}else{
			right = mid
		}
	}

	fmt.Println(left)
}

// 方程式を解く
func calc_equation(n float64, mid float64 )float64{
	saving :=  n + 1
	for i := 0;i < 5; i++ {
		saving = mid * saving + 1
	}
	return saving
}

