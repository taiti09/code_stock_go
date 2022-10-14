// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const INFINTY = 1000000000

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	left := 0
	right := len(lists)

	fmt.Println(lists)
	mergeSort(&lists, left, right)
	fmt.Println(lists)
}

func mergeSort(a *[]int, left int, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(a, left, mid)
		mergeSort(a, mid, right)
		merge(a, left, mid, right)
	}

}

func merge(a *[]int, left int, mid, right int) {
	L_len := (mid - left) + 1
	R_len := (right - mid) + 1

	var L []int
	var R []int

	L = make([]int, L_len, 10)
	R = make([]int, R_len, 10)

	for i := 0; i < L_len-1; i++ {
		L[i] = (*a)[left+i]
	}

	for i := 0; i < R_len-1; i++ {
		R[i] = (*a)[mid+i]
	}

	L[L_len-1] = INFINTY
	R[R_len-1] = INFINTY

	i := 0
	j := 0

	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			(*a)[k] = L[i]
			i++
		} else {
			(*a)[k] = R[j]
			j++
		}
	}
}