// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var lists = []int{4, 3, 7, 6, 5, 2, 1, 8}

	fmt.Println(lists)
	bubbleSort(&lists)
	fmt.Println(lists)
}

func bubbleSort(a *[]int) {
	for i := 0; i < len(*a)-1; i++ {
		for j := 1; j < len(*a)-i; j++ {
			if (*a)[j] < (*a)[j-1] {
				swap(a, j, j-1)
			}
		}
	}
}

func swap(a *[]int, i int, j int) {
	tmp := (*a)[i]
	(*a)[i] = (*a)[j]
	(*a)[j] = tmp
}
