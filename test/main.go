package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type book struct {
	ID        int
	Name      string
	Pressrun  int
	Publisher int
	Author    []int
}

type author struct {
	Name string
	ID   int
}

type publisher struct {
	Name string
	ID   int
}

func loadBooks(fileName string, x interface{}) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	dec.Decode(&x)
}
func loadPublishers(fileName string) []publisher {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	var temp []publisher
	dec.Decode(&temp)
	return temp
}

func main() {
	var books []book
	loadBooks("books.json",&books)

	var publishers []publisher
	loadBooks("publishers.json",&publishers)
	fmt.Println(books[0].Name + publishers[0].Name)

} 
