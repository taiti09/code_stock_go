package main

import (
	"fmt"
	"lib/sort"
	"sync"
)

func main() {
	var ans []int
	var list []int
	var n int
	var wg sync.WaitGroup
	var i int

	_, err := fmt.Scan(&n)
	sort.Checkerr(err)
	wg.Add(n)
	for i = 0 ; i < n ; i++ {
		var num int
		_, err := fmt.Scan(&num)
		list = append(list,num)
		sort.Checkerr(err)
	}

	for i = 0; i < n; i++ {
		go sort.Sleepsort(n,&ans,&wg)
	}

	wg.Wait()
	fmt.Println(ans)
}