// 入力例
// 8 2
// 3 1 4 1 5 9 2 6
// 5 3

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)
var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	inputs_a := make([]int, n)
	inputs_b := make([]int, m)
	for i := 0; i < n; i++ {
		inputs_a[i] = scanInt()
	}
	for i := 0; i < m; i++ {
		inputs_b[i] = scanInt()
	}

	// クイックソートで昇順にする
	sort.Slice(inputs_a, func(i, j int) bool {
		return inputs_a[i] < inputs_a[j]
	})

	for _, v := range inputs_b{
		// up_boundの二分探索
		left := -1
		right := n
		target := v
		for(1 < right - left){
			mid := left + (right - left) / 2
			if(target < inputs_a[mid]){
				right = mid
			}else{
				left = mid
			}
		}
		fmt.Println(left + 1)
	}
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	return i
}

