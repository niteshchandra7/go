// structs
// comparision of different objects is allowed for structs in Go!
// assignment creats a new object for structs!
package main

import "fmt"

type Book struct {
	title  string
	author string
	year   int
}

func (b *Book) updateAuthor(name string) {
	b.author = name
}

func main() {
	book1 := Book{
		title:  "The Devine Comedy",
		author: "DanteAligheri",
		year:   1320,
	}

	book2 := book1
	fmt.Printf("%p\n", &book1)
	fmt.Printf("%p\n", &book2)
	fmt.Printf("%T:%#v\n", book1, book1)
	book1.updateAuthor("ShakesPeare")
	fmt.Printf("%T:%#v\n", book1, book1)
	fmt.Printf("%T:%#v\n", book2, book2)
}
