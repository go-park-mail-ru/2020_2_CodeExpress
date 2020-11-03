package user

import "github.com/go-park-mail-ru/2020_2_CodeExpress/internal/models"

type UserRep interface {
	Insert(user *models.User) error
	Update(user *models.User) error
	UpdatePassword(user *models.User) error
	SelectByEmail(email string) (*models.User, error)
	SelectByName(name string) (*models.User, error)
	SelectByNameOrEmail(name string, email string, id uint64) (*models.User, error)
	SelectById(userID uint64) (*models.User, error)
	SelectWithPasswordById(id uint64) (*models.User, error)
	SelectWithPasswordByLogin(login string) (*models.User, error)
}
