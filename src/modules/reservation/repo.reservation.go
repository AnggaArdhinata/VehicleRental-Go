package reservation

import (
	"errors"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"gorm.io/gorm"
)

type reservation_repo struct {
	db *gorm.DB
}

func NewRepo(data *gorm.DB) *reservation_repo {
	return &reservation_repo{db: data}
}

func (r *reservation_repo) GetById(id uint) (*models.Reservation, error) {
	var data models.Reservation
	result := r.db.First(&data, "user_id = ?", id)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *reservation_repo) GetAll() (*models.Reservations, error) {
	var data models.Reservations
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *reservation_repo) Insert(data *models.Reservation) (*models.Reservation, error){
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *reservation_repo) Set(id uint, res *models.Reservation) (*models.Reservation, error) {
	result := r.db.Model(&res).Where("id = ?", id).Updates(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *reservation_repo) Remove(id uint) (*library.Messeage, error) {
	var data models.Reservation
	result := r.db.Delete(data, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library.Messeage{Msg: "successfully delete data"}, nil 
}


