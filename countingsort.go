// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const N = 8 //0～Nまでの整数の範囲でソートする

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	k := lists[len(lists)-1]
	fmt.Println(lists)
	result_lists := countingSort(&lists, k)

	fmt.Printf("[")
	for i := 1; i < len((*result_lists)); i++ {
		if i == len((*result_lists))-1 {
			fmt.Printf("%d", (*result_lists)[i])
			break
		}
		fmt.Printf("%d ", (*result_lists)[i])
	}
	fmt.Printf("]")
}

func countingSort(a *[]int, k int) *[]int {
	var b []int
	var c []int

	b = make([]int, N+1)
	c = make([]int, N+1)
	for i := 0; i <= N; i++ {
		c[i] = 0
	}

	for j := 0; j < len((*a)); j++ {
		c[(*a)[j]] = c[(*a)[j]] + 1
	}

	for i := 0; i < N; i++ {
		c[i+1] = c[i+1] + c[i]
	}

	for j := len((*a)) - 1; j >= 0; j-- {
		b[c[(*a)[j]]] = (*a)[j]
		c[(*a)[j]] = c[(*a)[j]] - 1
	}
	return &b
}
