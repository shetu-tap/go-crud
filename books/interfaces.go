package books

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type BookRepository interface {
	Create(book Book) *Book
	FindAll() []*Book
	Update(book *Book) *Book
	Delete(book Book)
	FindById(id uuid.UUID) *Book
}

type BookController interface {
	GetBooks(c *gin.Context)
	CreateBook(c *gin.Context)
	UpdateBook(c *gin.Context)
	GetDetail(c *gin.Context)
	DeleteBook(c *gin.Context)
}

type BookService interface {
	CreateBook(book Book) *Book
	FindBookByID(id uuid.UUID) *Book
	UpdateBook(id uuid.UUID, book Book) *Book
	DeleteBook(id uuid.UUID)
	FindAllBook() []*Book
}