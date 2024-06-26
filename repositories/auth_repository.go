package repositories

import (
	"errors"
	"gorm.io/gorm"
	"toyproject_recruiting_community/entities"
)

var UserNotFoundError = errors.New("[AuthRepository] user not found")
var RecordNotFoundError = errors.New("record not found")

type AuthRepository interface {
	FindById(id string) (*entities.User, error)
	Create(user *entities.User) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (ar *authRepository) Create(user *entities.User) error {
	tx := ar.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (ar *authRepository) FindById(id string) (*entities.User, error) {
	var user entities.User
	// Findメソッドを使うとエラーを吐き出さない、のでFirstに変更する
	tx := ar.db.First(&user, "id = ?", id)
	if tx.Error != nil {
		if tx.Error.Error() == RecordNotFoundError.Error() {
			return nil, UserNotFoundError
		}
		return nil, tx.Error
	}
	return &user, nil
}
