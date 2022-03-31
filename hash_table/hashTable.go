package hash_table

import (
	"fmt"
)

type HashTable struct {
	list [][]string
	size int
}

func (hashTable *HashTable) InitHashTable(size int) {
	hashTable.size = size
	hashTable.list = make([][]string, 10)
}

func (hashTable *HashTable) AppendElement(element string) {
	index := len(element) % hashTable.size
	hashTable.list[index] = append(hashTable.list[index], element)
}

func (hashTable *HashTable) PrintHashTable() {
	for i := 0; i < hashTable.size; i++ {
		fmt.Printf("%d: ", i)
		for _, element := range hashTable.list[i] {
			fmt.Printf("%s|", element)
		}
		fmt.Println()
	}
}
