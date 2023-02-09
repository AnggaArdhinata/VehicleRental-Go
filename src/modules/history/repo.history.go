package history

import (
	"errors"

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/library"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(data *gorm.DB) *history_repo {
	return &history_repo{db: data}
}

// func (r *history_repo) JoinData(idres uint) (*models.History, error) {
// 	var data models.History
// 	r.db.Preload("Reservation").Create(data)
// 	return &data, nil

// }

// func (r *history_repo) AddBooking(data *models.History) (*models.History, error) {
// 	var body models.Reservation
// 	if body.Isbooked == true {
// 		data.Reservation_Id = int(body.Id)
// 		data.Status = "on book"
// 		result := r.db.Create(&body)
// 		if result.Error != nil {
// 			return nil, errors.New("failed to get data")
// 		}
// 	}
// 	return data, nil
// }

func (r *history_repo) GetAll() (*models.Historis, error) {
	var data models.Historis
	result := r.db.Find(&data)
	if result.Error != nil {
		return nil, errors.New("failed to get data")
	}
	return &data, nil
}

func (r *history_repo) Insert(data *models.History) (*models.History, error) {
	result := r.db.Create(&data)
	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

func (r *history_repo) Set(id uint, res *models.History) (*models.History, error) {
	result := r.db.Model(&res).Where("id = ?", id).Updates(&res)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (r *history_repo) Remove(id uint) (*library.Messeage, error) {
	var data models.History
	result := r.db.Delete(data, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &library.Messeage{Msg: "successfully delete data"}, nil
}
