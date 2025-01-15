package main

import "fmt"


type HashTable struct{
	bucket [][]Entry
	size int
}

type Entry struct{
	key string
	value int
}

func (h *HashTable)Hash(key string) int {
	sum:=0;
	for _,v := range key{
	  sum += int(v)
	}
	return sum % h.size
}

func main()  {
	 h := &HashTable{
		bucket: make([][]Entry, 10),
		size: 10,
	 }

	 h.Insert("hello",100)
	 h.Insert("world",200)
	 h.Insert("globe",300)

	 fmt.Println(h)

}

//Insert - 
 //hash the key get index
 //loop thru H[index][i].key = key if yes update value  or else append a bucket to H[index][i]

func (h *HashTable)Insert(key string , value int)  {
	entry := &Entry{key: key,value: value}

	index := h.Hash(key)
	fmt.Printf("hash for %v is %v \n",key,index)

	for i,v := range h.bucket[index]{
		if v.key == key{
			h.bucket[index][i].value = value
			return
		}
	}

	h.bucket[index] = append(h.bucket[index], *entry)
}

//Delete -
  //go thru index slice remove it by skip and append

//Get - 
  //go thru search key return value 
