package controllers

import (
	response "golang-base-code/src/app/core"
	middleware "golang-base-code/src/app/middlewares"
	model "golang-base-code/src/app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type bookHandler struct {
	Repo middleware.BookMiddleware
}

func BookHandler(db *gorm.DB) *bookHandler {
	return &bookHandler{
		Repo: middleware.BookConnectionMw(db),
	}
}

func (b *bookHandler) GetAll(c echo.Context) error {
	book, err := b.Repo.Fetch()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success get all data", book))
}

func (b *bookHandler) GetById(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter("failed get detail book, please insert number on param", nil))
	}

	book, err := b.Repo.GetById(int32(bookId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success get detail book", book))
}

func (b *bookHandler) Create(c echo.Context) (err error) {
	payload := new(model.Book)
	if err = c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	book, err := b.Repo.Create(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success create new book", book))
}

func (b *bookHandler) Update(c echo.Context) (err error) {
	payload := new(model.Book)
	if err = c.Bind(payload); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	if err = c.Validate(payload); err != nil {
		c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	book, err := b.Repo.Update(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success update book", book))
}

func (b *bookHandler) Delete(c echo.Context) error {
	bookId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter("failed delete detail book, please insert number on param", nil))
	}

	book, err := b.Repo.Delete(int32(bookId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseFormatter(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.ResponseFormatter("success delete book", book))
}
