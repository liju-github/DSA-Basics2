package main

import "fmt"

type LinkedList struct{
	start *node
}

type node struct{
	value int
	next *node
}

func main(){
	l := &LinkedList{}

	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Display()
	fmt.Println("last element is : ",l.Peek())
    l.Pop()
	l.Display()
}

func (l *LinkedList)Push(value int){
	newnode:=&node{
		value: value,
	}
	
	if l.start == nil{
		l.start = newnode
	    return
	}
	
	current := l.start
	for current.next!=nil{
		current = current.next
	}
	
	current.next = newnode
}

func (l *LinkedList)Display(){
	current := l.start
	for current!=nil{
		fmt.Println(current.value)
		current = current.next
	}
}

func (l *LinkedList)Peek() int {
	current := l.start
	for current.next != nil{
		current = current.next
	}
	
	return current.value
}

func (l *LinkedList)Pop() {
	current := l.start
	if current == nil {
		fmt.Println("stack is empty")
		return
	}
	if current.next==nil{
		l.start = nil
		return
	}
	for current.next.next != nil{
		current = current.next
	}
	
	current.next = nil
}