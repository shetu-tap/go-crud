package books

import uuid "github.com/satori/go.uuid"

type bookService struct {
	repository BookRepository
}

func (b *bookService) DeleteBook(id uuid.UUID) {
	book := b.repository.FindById(id)
	b.repository.Delete(*book)
}

func (b *bookService) CreateBook(book Book) *Book {
	return b.repository.Create(book)
}

func (b *bookService) FindBookByID(id uuid.UUID) *Book {
	book := b.repository.FindById(id)
	return book
}

func (b *bookService) UpdateBook(id uuid.UUID, book Book) *Book {
	bookInstance := b.repository.FindById(id)
	bookInstance.Author = book.Author
	bookInstance.Description = book.Description
	bookInstance.Name = book.Name
	bookInstance = b.repository.Update(bookInstance)
	return bookInstance
}


func (b *bookService) FindAllBook() []*Book {
	return b.repository.FindAll()
}

func NewBookService(repository BookRepository) BookService {
	return &bookService{
		repository: repository,
	}
}
