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
	m := scanInt()
	inputs_a := make([]int, n)
	inputs_b := make([]int, m)
	for i := 0; i < n; i++ {
		inputs_a[i] = scanInt()
	}
	for i := 0; i < m; i++ {
		inputs_b[i] = scanInt()
	}

	for _, v := range inputs_b{
		// lower_boundの二分探索
		left := -1
		right := n
		target := v
		for(1 < right - left){
			mid := left + (right - left) / 2
			if(target <= inputs_a[mid]){
				right = mid
			}else{
				left = mid
			}
		}
		fmt.Println(right)
	}
}
