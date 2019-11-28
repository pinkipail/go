package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	NumOfBooks = 5
)

type Book struct {
	ID        int
	Name      string
	Pressrun  int
	Publisher int
	Author    []int
}

type Author struct {
	Name string
	ID   int
}

type Publisher struct {
	Name string
	ID   int
}

func loadJson(fileName string, x interface{}) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	dec.Decode(&x)
}

func (author Author) WriteBook(publisher1 chan<- Book, publisher2 chan<- Book) {
	for i := 0; i < NumOfBooks; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		var book Book
		book.Author = append(book.Author, author.ID)
		fmt.Printf("Автор %v написал книгу\n", author.Name)
		select {
		case publisher1 <- book:
			fmt.Printf("Автор %v направил книгу в первое издательство\n", author.Name)
		case publisher2 <- book:
			fmt.Printf("Автор %v направил книгу во второе издательство\n", author.Name)

		}
	}
}

func (publisher Publisher) PublishBook(c <-chan Book, books *[]Book, authors []Author) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		book := <-c
		book.Publisher = publisher.ID
		book.Pressrun = rand.Intn(10000)
		*books = append(*books, book)
		fmt.Printf("Издательство %v напечатало книгу автора %v\n", publisher.Name, authors[book.Author[0] - 1].Name)
	}
}

func main() {

	var publishers []Publisher
	loadJson("publishers.json",&publishers)

	var authors []Author
	loadJson("author.json",&authors)

	publisher1 := make(chan Book, 1)
	publisher2 := make(chan Book, 1)
	rand.Seed(time.Now().UnixNano())

	var t string
	var books []Book

	for _, author := range authors {
		go author.WriteBook(publisher1, publisher2)
	}


	go publishers[0].PublishBook(publisher1, &books,authors)
	go publishers[1].PublishBook(publisher2, &books,authors)
	fmt.Scanln(&t)

	for i := 0; i < 6; i++ {
		var pc1, pc2 float32 = 0, 0
		for _, b := range books {
			if authors[i].ID == b.Author[0] {
				if b.Publisher == publishers[0].ID {
					pc1++
				} else {
					pc2++
				}
			}
		}

		fmt.Printf("%v: Перое издательство - %.0f процентов. Второе издательство - %.0f процентов.\n", authors[i].Name, float32(pc1/NumOfBooks*100), float32(pc2/NumOfBooks*100))
	}
}
