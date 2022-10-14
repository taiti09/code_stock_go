// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const INFINTY = 1000000000

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	p := 0
	r := len(lists) - 1

	fmt.Println(lists)
	quickSort(&lists, p, r)
	fmt.Println(lists)
}

func quickSort(a *[]int, p int, r int) {
	if p < r {
		q := parttion(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q+1, r)
	}
}

func parttion(a *[]int, p int, r int) int {
	x := (*a)[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if (*a)[j] <= x {
			i++
			swap(a, i, j)
		}
	}
	swap(a, i+1, r)
	return i + 1
}

func swap(a *[]int, i int, j int) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}
