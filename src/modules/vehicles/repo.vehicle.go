package vehicles

import (
	"errors"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"gorm.io/gorm"
)

type vehicles_repo struct {
	db *gorm.DB
}

func NewRepo(data *gorm.DB) *vehicles_repo {
	return &vehicles_repo{db: data}
}

func (r *vehicles_repo) GetAll() (*models.Vehicles, error) {
	var data models.Vehicles
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *vehicles_repo) Insert(data *models.Vehicle) (*models.Vehicle, error){
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *vehicles_repo) Set(id uint, res *models.Vehicle) (*models.Vehicle, error) {
	result := r.db.Model(&res).Where("id = ?", id).Updates(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *vehicles_repo) Remove(id uint) (*library.Messeage, error) {
	var data models.Vehicle
	result := r.db.Delete(data, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library.Messeage{Msg: "successfully delete data"}, nil 
}


