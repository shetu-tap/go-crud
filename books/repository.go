package books

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type bookRepository struct {
	DB *gorm.DB
}

func (u *bookRepository) FindById(id uuid.UUID) *Book {
	var book Book
	u.DB.First(&book, id)
	return &book
}

func (u *bookRepository) Create(book Book) *Book {
	u.DB.Create(&book)
	return &book
}

func (u *bookRepository) FindAll() []*Book {
	var books []*Book
	u.DB.Find(&books)
	return books
}

func (u *bookRepository) Update(book *Book) *Book {
	u.DB.Save(&book)
	return book
}

func (u *bookRepository) Delete(book Book) {
	u.DB.Delete(&book)
	return
}

func NewRepository() BookRepository {
	return &bookRepository{
		DB: ConnectToDB(),
	}
}
