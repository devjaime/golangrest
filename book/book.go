package book

import(
	"github/gofiber/fiber"
)

func GetBooks(c *fiber.Ctx){
	c.Send("All Books")
}


func GetBook(c *fiber.Ctx){
	c.Send("A Single Book")
}


func NewBook(c *fiber.Ctx){
	c.Send("Adds a new book")
}


func DeleteBook(c *fiber.Ctx){
	c.Send("Deletes a book")
}