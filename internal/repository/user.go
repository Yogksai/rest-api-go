package repository

import (
	"github.com/Yogksai/rest-api-go/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

//Конструктор
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}


func (r *UserRepo) Create(user *models.User) error {
	err := r.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepo) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) Update(user *models.User) error {
	err := r.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) Delete(id uint) error {
	err := r.DB.Delete(&models.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}


