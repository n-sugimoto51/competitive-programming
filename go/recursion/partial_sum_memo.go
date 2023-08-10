// 部分和問題をメモ化再帰で解く
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
// panic: runtime error: index out of range [-3]

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	x := scanInt()// この値を作ることができるかどうかを判定する
	inputs_a := make([]int, n)
	for i := 0; i < n; i++ {
		inputs_a[i] = scanInt()
	}
	// momoスライスの要素全てを-1で初期化
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo = append(memo, make([]int, x+1))
		memo[i] = make([]int, x+1)
		for j := 0; j < x+1; j++ {
			memo[i][j] = -1
		}
	}

	// memoの中身を確認
	fmt.Println(memo)

	// 再帰で全探索
	if(calc_partial_sum_memo(n, inputs_a, x, memo) != 0){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}

func calc_partial_sum_memo(n int, inputs_a []int, x int, memo [][]int) int {
	if n == 0 {
		if x == 0 {
			return 1
		} else {
			return 0
		}
	}
	if memo[n][x] != -1 {
		return memo[n][x]
	}

	if calc_partial_sum_memo(n-1, inputs_a, x, memo) == 1 {
		memo[n][x] = 1
		return 1
	}

	if calc_partial_sum_memo(n-1, inputs_a, x-inputs_a[n-1], memo) == 1 {
		memo[n][x] = 1
		return 1
	}

	memo[n][x] = 0
	return 0
}
