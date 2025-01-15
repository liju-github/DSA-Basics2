package main

import "fmt"

type HashTable struct {
	bucket [][]Entry
	size   int
}

type Entry struct {
	key   int
	value int
}

func (ht *HashTable) NewHashTable(size int) *HashTable {
	bucket := make([][]Entry, size)
	return &HashTable{
		bucket: bucket,
		size:   size,
	}
}

func (ht *HashTable) Hash(key int) int {
	return key % ht.size
}

func (ht *HashTable) Insert(key, value int) {
	index := ht.Hash(key)
	for i, entry := range ht.bucket[index] {
		if entry.key == key {
			ht.bucket[index][i].value = value
			return
		}
	}
	ht.bucket[index] = append(ht.bucket[index], Entry{key: key, value: value})
}

func (ht *HashTable) Delete(key int) {
	index := ht.Hash(key)
	for i, entry := range ht.bucket[index] {
		if entry.key == key {
			ht.bucket[index] = append(ht.bucket[index][:i], ht.bucket[index][i+1:]...)
			return
		}
	}
}

func (ht *HashTable) Search(key int) (int, bool) {
	index := ht.Hash(key)
	for _, entry := range ht.bucket[index] {
		if entry.key == key {
			return entry.value, true
		}
	}
	return 0, false
}

func main() {
	ht := &HashTable{}
	ht = ht.NewHashTable(10)

	ht.Insert(1, 10)
	ht.Insert(2, 20)
	ht.Insert(2, 30)

	ht.Insert(3, 30)
	fmt.Println(ht.Search(3))

	fmt.Println(ht)
}
