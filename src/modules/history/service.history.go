package history

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type history_service struct {
	repo interfaces.History_RepoIF
}

func NewService(repo interfaces.History_RepoIF) *history_service {
	return &history_service{repo}
}

// func (s *history_service) JoinReserv(idres uint) *library.Responses {
// 	data, err := s.repo.JoinData(idres)
// 		if err != nil {
// 			return library.Response(err.Error(), 400, true)
// 		}
// 	return library.Response(data, 200, false)
// }

// func (s *history_service) AddFromReserv(data *models.History) *library.Responses {
// 	data, err := s.repo.AddBooking(data)
// 		if err != nil {
// 			return library.Response(err.Error(), 400, true)
// 		}
// 	return library.Response(data, 200, false)

// }

func (s *history_service) FindAll() *library.Responses {
	data, err := s.repo.GetAll()
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *history_service) Add(data *models.History) *library.Responses {
	data, err := s.repo.Insert(data)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *history_service) Update(id uint, res *models.History) *library.Responses {
	data, err := s.repo.Set(id, res)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *history_service) Delete(id uint) *library.Responses {
	data, err := s.repo.Remove(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}