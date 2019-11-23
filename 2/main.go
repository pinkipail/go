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

func loadBooks(fileName string) []book {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	var temp []book
	dec.Decode(&temp)
	return temp
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

func loadAuthor(fileName string) []author {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)

	var temp []author
	dec.Decode(&temp)
	return temp
}

func countPressrun(authors []author, books []book) {
	var Select int
	for _, author := range authors {
		fmt.Println(author.ID, " - ", author.Name)
	}
	fmt.Println("Введите номер автора:")
	var sumPressrun int
	fmt.Fscan(os.Stdin, &Select)
	for _, book := range books {
		for _, int := range book.Author {
			if Select == int {
				sumPressrun += book.Pressrun
			}
		}
	}
	fmt.Println("Общий тираж книг автора:", sumPressrun)

}

func authorsInPublishing(authors []author, books []book, publishers []publisher) {
	for _, publisher := range publishers {
		fmt.Println(publisher.Name, " - ", publisher.ID)
	}
	var Select int

	authorArr := []int{}
	fmt.Println("Введите номер издательства:")

	fmt.Fscan(os.Stdin, &Select)
	for _, book := range books {
		if Select == book.Publisher {
			for _, int := range book.Author {
				authorArr = append(authorArr, int)
			}
		}
	}
	unique := map[int]string{}

	for _, v := range authorArr {
		unique[v] = authors[v-1].Name
	}

	fmt.Println("Авторы с которые публиковались в издательстве")
	for _, str := range unique {
		fmt.Println(str)
	}

}

func bestAuthor(books []book, authors []author) {
	for _, book := range books {
		fmt.Println(book.ID, " - ", book.Name)
	}
	var Select int

	arrCount := []int{}

	fmt.Println("Введите номер книги:")
	fmt.Fscan(os.Stdin, &Select)
	for _, book := range books {
		if book.ID == Select {
			for _, int := range book.Author {
				arrCount = append(arrCount, int)
			}

		}
	}
	fmt.Println(arrCount)
	lenArr := []int{0, 0, 0, 0, 0}
	for _, book := range books {
		for j, int := range book.Author {
			for i := 0; i < len(arrCount); i++ {
				if arrCount[i] == int {
					lenArr[i]++
				}

			}
			fmt.Println(j)
		}
	}
	max := 0
	for i := 0; i < len(lenArr); i++ {
		if lenArr[i] > max {
			max = lenArr[i]
		}
	}
	fmt.Println(books[Select-1].Name, " - ", authors[max-1].Name)

}
func main() {

	books := loadBooks("books.json")

	publishers := loadPublishers("publishers.json")

	authors := loadAuthor("author.json")
	countPressrun(authors, books)
	fmt.Println()

	authorsInPublishing(authors, books, publishers)
	fmt.Println()
	bestAuthor(books, authors)

}
