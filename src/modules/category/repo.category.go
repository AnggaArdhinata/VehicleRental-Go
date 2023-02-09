package category

import (
	"errors"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"gorm.io/gorm"
)

type category_repo struct {
	db *gorm.DB
}

func NewRepo(data *gorm.DB) *category_repo {
	return &category_repo{db: data}
}

func (r *category_repo) GetById(id uint) (*models.Category, error) {
	var data models.Category
	result := r.db.First(&data, "id = ?", id)
	r.db.Preload("Vehicles").Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *category_repo) GetAll() (*models.Categories, error) {
	var data models.Categories
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	if len(data) <= 0 {
		return nil, errors.New("data category is empty !")
	}
	return &data, nil
}

func (r *category_repo) Insert(data *models.Category) (*models.Category, error) {
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *category_repo) Set(id uint, res *models.Category) (*models.Category, error) {
	result := r.db.Model(&res).Where("id = ?", id).Updates(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *category_repo) Remove(id uint) (*library.Messeage, error) {
	var data models.Category
	result := r.db.Delete(data, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library.Messeage{Msg: "successfully delete data"}, nil
}
