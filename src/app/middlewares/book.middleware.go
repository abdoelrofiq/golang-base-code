package domain

import (
	"errors"
	"fmt"
	model "golang-base-code/src/app/models"
	repository "golang-base-code/src/http/repository/books"

	"gorm.io/gorm"
)

var books []model.Books
var book model.Books

type mysqlBooksDomain struct {
	Db *gorm.DB
}

func MysqlBooksDomain(connection *gorm.DB) repository.BooksRepo {
	return &mysqlBooksDomain{
		Db: connection,
	}
}

func (m *mysqlBooksDomain) Fetch() ([]model.Books, error) {
	m.Db.Raw("SELECT * FROM books").Scan(&books)

	return books, nil
}

func (m *mysqlBooksDomain) GetById(bookId int32) (model.Books, error) {
	m.Db.Raw("SELECT * FROM books WHERE id = ?", bookId).Scan(&book)
	if book.Id == 0 {
		return book, errors.New("Book not found")
	}

	fmt.Println(bookId)
	fmt.Println(book)

	return book, nil
}

func (m *mysqlBooksDomain) Create(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Create(book)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return *book, nil
}

func (m *mysqlBooksDomain) Update(book *model.Books) (model.Books, error) {
	result := m.Db.Model(&book).Update("id", book.Id)
	if result.Error != nil {
		return model.Books{}, result.Error
	}

	return *book, nil
}

func (m *mysqlBooksDomain) Delete(bookId int32) (bool, error) {
	result := m.Db.Delete(&model.Books{}, bookId)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
