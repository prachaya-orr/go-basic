package main

import (
	"encoding/json"
	"fmt"

	"github.com/prachaya-orr/go-basic/models"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(&obj, "", "\t")
	fmt.Println(string(raw))
}

type IBook interface {
	GetBook() *book
	SetTitle(title string)
}

type book struct {
	*models.Book
}

func NewBook(title, author string) IBook {
	return &book{
		&models.Book{
			Id:     1,
			Title:  title,
			Author: author,
		},
	}
}

func (b *book) GetBook() *book {
	return b
}

func (b *book) SetTitle(title string) {
	b.Title = title
}

func main() {
	b := NewBook("Golang Basic", "Gopher")
	fmt.Println(b.GetBook())
	raw, _ := json.MarshalIndent(&b, "", "\t")
	fmt.Println(string(raw))
	b.SetTitle("ABC")
	Debug(b.GetBook())

}
