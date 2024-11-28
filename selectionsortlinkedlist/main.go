package main

import "fmt"

type node struct {
	value int
	next *node
}

type linkedlist struct{
	head *node
}

func (l *linkedlist)selectionSortLinkedList() {
	if l.head==nil {
		return 
	}

	for sorted:=l.head;sorted!=nil;sorted=sorted.next{
		minnode := sorted
		for current := sorted.next; current != nil; current = current.next {
			if current.value < minnode.value {
				minnode = current
			}
		}
		sorted.value, minnode.value = minnode.value, sorted.value
	}
}

func (l *linkedlist)Print() {
	for current := l.head; current != nil; current = current.next {
		fmt.Printf("%d ", current.value)
	}
	fmt.Println()
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
	l.selectionSortLinkedList()
	fmt.Println("sorted  linkedlist := ")
	l.Print()
}
