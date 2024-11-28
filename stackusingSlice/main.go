package main

import "fmt"


type Stack struct{
	items []int
}


func (s *Stack)Push(value int)  {
	s.items = append(s.items,value)
}


func (s *Stack)Pop() int {
	l := len(s.items)
	lastElement := s.items[l-1]
	s.items = s.items[:l-1]
	
	return lastElement
}

func (s *Stack)Peek()int  {
	l := len(s.items)
	return s.items[l-1]
}

func main()  {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.items)
	fmt.Println("last element is: ",s.Peek())
	fmt.Println("last element to pop is: ",s.Pop())
	fmt.Println(s.items)

}