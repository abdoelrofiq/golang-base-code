package middlewares

import (
	"errors"
	model "golang-base-code/src/app/models"

	"gorm.io/gorm"
)

var books []model.Book
var book model.Book

type BookMiddleware interface {
	Fetch() ([]model.Book, error)
	GetById(id int32) (model.Book, error)
	Create(b *model.Book) (model.Book, error)
	Update(b *model.Book) (model.Book, error)
	Delete(id int32) (bool, error)
}

type bookMiddlewareBuilder struct {
	DB *gorm.DB
}

func BookConnectionMw(connection *gorm.DB) BookMiddleware {
	return &bookMiddlewareBuilder{
		DB: connection,
	}
}

func (m *bookMiddlewareBuilder) Fetch() ([]model.Book, error) {
	m.DB.Find(&books)

	return books, nil
}

func (m *bookMiddlewareBuilder) GetById(bookId int32) (model.Book, error) {
	m.DB.Where("id = ?", bookId).Find(&book)
	if book.Id == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (m *bookMiddlewareBuilder) Create(book *model.Book) (model.Book, error) {
	result := m.DB.Model(&book).Create(book)
	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Update(book *model.Book) (model.Book, error) {
	result := m.DB.Model(&book).Update("id", book.Id)
	if result.Error != nil {
		return model.Book{}, result.Error
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Delete(bookId int32) (bool, error) {
	result := m.DB.Delete(&model.Book{}, bookId)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
