package main

import "fmt"

type linkedlist struct {
	head *node
}

type node struct {
	value int
	next  *node
}

func main() {
	l := &linkedlist{}

	l.head = &node{value: 5}
	l.head.next = &node{value: 8}
	l.head.next.next = &node{value: 1}
	l.head.next.next.next = &node{value: 2}
	l.head.next.next.next.next = &node{value: 3}
	fmt.Println("unsorted  linkedlist := ")
	l.Print()
	l.Bubblesort()
	fmt.Println("sorted  linkedlist := ")
	l.Print()

}

func (l *linkedlist) Bubblesort() {
	if l.head == nil {
		return
	}
	swapped := true
	for swapped {
		swapped = false
		current := l.head

		for current.next != nil {
			if current.value > current.next.value {
				current.value, current.next.value = current.next.value, current.value
				swapped = true
			}
			current = current.next
		}
	}
}

func (l *linkedlist) Print() {
	current := l.head

	for current != nil {
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Print("\n")
}
