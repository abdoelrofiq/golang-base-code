package middlewares

import (
	"errors"
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

var books []model.Books
var book model.Books

type BooksMiddleware interface {
	Fetch() ([]model.Books, error)
	GetById(id int32) (model.Books, error)
	Create(b *model.Books) (model.Books, error)
	Update(b *model.Books) (model.Books, error)
	Delete(id int32) (bool, error)
}

type bookMiddlewareBuilder struct {
	Db *gorm.DB
}

func BookConnectionMw(connection *gorm.DB) BooksMiddleware {
	return &bookMiddlewareBuilder{
		Db: connection,
	}
}

func (m *bookMiddlewareBuilder) Fetch() ([]model.Books, error) {
	m.Db.Find(&books)

	return books, nil
}

func (m *bookMiddlewareBuilder) GetById(bookId int32) (model.Books, error) {
	m.Db.Where("id = ?", bookId).Find(&book)
	if book.Id == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (m *bookMiddlewareBuilder) Create(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Create(book)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Update(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Update("id", book.Id)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Delete(bookId int32) (bool, error) {
	result := m.Db.Delete(&model.Books{}, bookId)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
