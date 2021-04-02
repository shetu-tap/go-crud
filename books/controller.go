package books

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type bookController struct {
	service BookService
}

func (b *bookController) GetBooks(c *gin.Context) {
	books := b.service.FindAllBook()
	c.JSON(http.StatusOK, books)
}

func (b *bookController) CreateBook(c *gin.Context) {
	log.Println("Creating a book")
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//c.BindJSON(&book)
	bookInstance := b.service.CreateBook(book)
	c.JSON(http.StatusOK, bookInstance)
}

func (b *bookController) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var bookInstance Book
	err = c.ShouldBindJSON(&bookInstance)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	book := b.service.UpdateBook(uid, bookInstance)
	c.JSON(http.StatusOK, book)
}

func (b *bookController) GetDetail(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	book := b.service.FindBookByID(uid)
	if book.Name == "" {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, book)
}

func (b *bookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	uid, err := uuid.FromString(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	b.service.DeleteBook(uid)
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully",
	})
}

func NewBookController(service BookService) BookController {
	return &bookController{
		service: service,
	}
}
