package books

import (
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID uuid.UUID `gorm:"type:UUID;primary_key"json:"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	Description string `json:"description"`
}

func (book *Book) BeforeCreate(tx *gorm.DB) (err error) {
	uid := uuid.NewV4()
	book.ID = uid
	return
}
