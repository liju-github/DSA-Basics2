package main

import "fmt"



type Queue struct{
	elements []int
}


func main()  {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	q.Display()
	q.Dequeue()
	q.Display()

}


func (q *Queue)Display()  {
	fmt.Println(q.elements)
}

func (q *Queue)Enqueue(value int)  {
	q.elements = append(q.elements, value)
}


func (q *Queue)Dequeue() int {
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}