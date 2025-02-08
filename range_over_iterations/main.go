package main

import "fmt"

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next  *element[T]
	value T
}

func (lst *List[T]) push(value T) {
	if lst.tail == nil {
		lst.head = &element[T]{value: value}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{value: value}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) getAll() {
	i := lst.head
	for i != nil {
		fmt.Println(i.value)
		i = i.next
	}
}

func main() {
	lst := &List[int]{}
	lst.push(19)
	lst.push(20)
	lst.push(7)
	lst.push(2)
	lst.getAll()
}
