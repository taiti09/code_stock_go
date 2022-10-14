// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const N = 8 //0～Nまでの整数の範囲でソートする

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	fmt.Println(lists)
	heapSort(&lists)
	fmt.Println(lists)
}

func heapSort(a *[]int) {
	for i := len((*a)) / 2; i >= 0; i-- {
		downheap(a, i, len((*a))-1)
	}

	for i := len((*a)) - 1; i >= 1; i-- {
		swap(a, 0, i)
		downheap(a, 0, i-1)
	}

}

func downheap(a *[]int, i int, n int) {
	var j int

	for {
		j = 2*i + 1
		if j > n {
			break
		}
		if j < n && (*a)[j+1] > (*a)[j] {
			j++
		}
		if (*a)[i] >= (*a)[j] {
			break
		}
		swap(a, i, j)
		i = j
	}
}

func swap(a *[]int, i int, j int) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}
