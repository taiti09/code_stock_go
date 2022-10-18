// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Queue struct {
	data []int
	size int
}

type Node struct {
	dist []int
	seen bool
}

func main() {
	var G map[int]*Node
}

func BFS(G *map[int]*Node, v int) {
	var Q *Queue

	Q = NewQueue(10)

	(*G)[v].seen = true
	Q.push(v)

	for !Q.isEmpty() {
		v = Q.pop()
		for _, i := range (*G)[v].dist {
			if (*G)[i].seen == false {
				(*G)[i].seen = true
				Q.push(i)
			}
		}
	}
}

func NewQueue(cap int) *Queue {
	return &Queue{
		data: make([]int, 0, cap),
		size: 0,
	}
}

func (q *Queue) push(p int) {
	q.data = append(q.data, p)
	q.size++
}

func (q *Queue) pop() int {
	if q.isEmpty() {
		return -1
	}
	res := q.data[0]
	q.data = q.data[1:]
	q.size = q.size - 1

	return res
}

func (q *Queue) isEmpty() bool {
	if q.size == 0 {
		return true
	}

	return false
}

func (q *Queue) front() int {
	if q.isEmpty() {
		return -1
	}
	return q.data[0]
}

func (q *Queue) string() string {
	return fmt.Sprint(q.data)
}