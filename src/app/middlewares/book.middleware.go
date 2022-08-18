package middlewares

import (
	"errors"
	"golang-base-code/src/app/core"
	model "golang-base-code/src/app/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var books []model.Book
var book model.Book

type BookMiddleware interface {
	Fetch(c echo.Context) ([]model.Book, error)
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

func (m *bookMiddlewareBuilder) Fetch(c echo.Context) ([]model.Book, error) {
	FQP, err := core.FQP(m.DB, c)
	if err != nil {
		return books, errors.New(err.Error())
	}

	result := FQP.Find(&books)
	if result.Error != nil {
		return books, errors.New(result.Error.Error())
	}

	return books, nil
}

func (m *bookMiddlewareBuilder) GetById(bookId int32) (model.Book, error) {
	result := m.DB.Where("id = ?", bookId).Find(&book)
	if result.Error != nil {
		return book, errors.New(result.Error.Error())
	}

	if book.Id == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}

func (m *bookMiddlewareBuilder) Create(book *model.Book) (model.Book, error) {
	result := m.DB.Model(&book).Create(book)
	if result.Error != nil {
		return model.Book{}, errors.New(result.Error.Error())
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Update(book *model.Book) (model.Book, error) {
	result := m.DB.Model(&book).Where("id = ?", book.Id).Updates(book)
	if result.Error != nil {
		return model.Book{}, errors.New(result.Error.Error())
	}

	return *book, nil
}

func (m *bookMiddlewareBuilder) Delete(bookId int32) (bool, error) {
	result := m.DB.Delete(&model.Book{}, bookId)
	if result.Error != nil {
		return false, errors.New(result.Error.Error())
	}

	return true, nil
}
