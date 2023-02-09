package users

import (
	"errors"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"gorm.io/gorm"
)

type users_repo struct {
	db *gorm.DB
}

func NewUserRepo(data *gorm.DB) *users_repo {
	return &users_repo{db: data}
}

func (r *users_repo) GetById(id uint) (*models.User, error) {
	var data models.User
	result := r.db.First(&data, "id = ?", id)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	data.Password = "secret"
	return &data, nil
}

func (r *users_repo) FindByEmail(email string) (*models.User, error) {
	var data models.User
	result := r.db.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *users_repo) FindByName(username string) (*models.User, error) {
	var data models.User
	result := r.db.First(&data, "name = ?", username)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}


func (r *users_repo) GetAll() (*models.Users, error) {
	var data models.Users
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	if len(data) <= 0 {
		return nil, errors.New("data users is empty !")
	}
	for i := range data {
		data[i].Password = "secret"
	}
	return &data, nil
}

func (r *users_repo) Insert(data *models.User) (*models.User, error) {
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *users_repo) Set(id uint, res *models.User) (*models.User, error) {
	result := r.db.Model(&res).Where("id = ?", id).Updates(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *users_repo) Remove(id uint) (*library.Messeage, error) {
	var data models.User
	result := r.db.Delete(data, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library.Messeage{Msg: "successfully delete data"}, nil
}
