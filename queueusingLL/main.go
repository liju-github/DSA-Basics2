package main

import "fmt"

type linkedlist struct{
    head *node
}

type node struct{
	value int
	next *node
}

func main()  {
	l := linkedlist{}
	l.Enqueue(1)
	l.Enqueue(2)
	l.Enqueue(3)
	l.Enqueue(4)
	l.Enqueue(5)
	l.Display()
	fmt.Println(l.Dequeue())
	l.Display()
}

func (l *linkedlist)Enqueue(value int)  {
	current := l.head

	if current == nil{
		l.head = &node{value: value}
		return
	}

	for current.next!=nil{
      current= current.next
	}

	newnode:=&node{value: value}
	current.next = newnode

}

func (l *linkedlist)Dequeue() int {
	current := l.head
	if current == nil{
		return 0
	}

	value := current.value
	l.head = l.head.next
	return value
}



func (l *linkedlist)Display()  {
	current := l.head
	if current == nil{
		fmt.Println("queue is empty")
		return
	}

	for current !=nil{
        fmt.Printf("%v ",current.value)
		current=current.next
	}

	fmt.Println()

}