// You can edit this code!
// Click here and start typing.
package main

import "fmt"

const INFINTY = 1000000000

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	fmt.Println(lists)
	selectionSort(&lists)
	fmt.Println(lists)
}

func selectionSort(a *[]int) {
	var min int
	for i := 0; i < len((*a)); i++ {
		min = i
		for j := i; j < len((*a)); j++ {
			if (*a)[j] < (*a)[min] {
				min = j
			}
		}
		swap(a, i, min)
	}
}

func swap(a *[]int, i int, j int) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}
