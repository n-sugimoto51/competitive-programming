// 部分和問題
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	x := scanInt()// この値を作ることができるかどうかを判定する
	inputs_a := make([]int, n)
	for i := 0; i < n; i++ {
		inputs_a[i] = scanInt()
	}

	// 再帰で全探索
	if(calc_partial_sum(n, inputs_a, x)){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}

func calc_partial_sum(n int, inputs_a []int, x int) bool {
	if n == 0 {
		if x == 0 {
			return true
		} else {
			return false
		}
	}

	if calc_partial_sum(n-1, inputs_a, x) {
		return true
	}

	if calc_partial_sum(n-1, inputs_a, x-inputs_a[n-1]) {
		return true
	}

	return false
}
