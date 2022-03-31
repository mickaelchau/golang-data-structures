package main

import (
	"awesomeProject/hash_table"
	"awesomeProject/tree"
	"fmt"
)

func main() {
	myTree := tree.InitTree()
	myTree.Dfs()
	fmt.Println("End of the program")
	var hashTable hash_table.HashTable
	hashTable.InitHashTable(10)
	hashTable.AppendElement("mange")
	hashTable.AppendElement("Totalité")
	hashTable.AppendElement("Compréhension")
	hashTable.AppendElement("Soir je pense que ça va être chaud")
	hashTable.AppendElement("ce")
	hashTable.AppendElement("Gucci")
	hashTable.AppendElement("je mange")
	hashTable.PrintHashTable()
}
