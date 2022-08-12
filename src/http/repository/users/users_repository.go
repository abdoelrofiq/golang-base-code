package repository

import model "golang-base-code/src/app/models"

type UserRepo interface {
	Fetch() ([]model.User, error)
}
