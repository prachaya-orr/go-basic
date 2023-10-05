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

type book struct {
	*models.Book
}

func (b *book) GetBook() *book {
	return b
}

func (b *book) SetTitle(title string) {
	b.Title = title
}

func main() {
	b := &book{
		&models.Book{
			Id:     1,
			Title:  "Golang Basic",
			Author: "arther",
		},
	}
	b.SetTitle("ABC")
	Debug(b.GetBook())

}
