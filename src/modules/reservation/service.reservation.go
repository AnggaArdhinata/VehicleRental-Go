package reservation

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type reservation_service struct {
	repo interfaces.Reservation_RepoIF
	
}

// type history struct {
// 	history interfaces.Reservation_RepoIF
// }

func NewService(repo interfaces.Reservation_RepoIF) *reservation_service {
	return &reservation_service{repo}
}

func (s *reservation_service) FindByIdUser(id uint) *library.Responses {
	data, err := s.repo.GetById(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *reservation_service) FindAll() *library.Responses {
	data, err := s.repo.GetAll()
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *reservation_service) Add(data *models.Reservation) *library.Responses {
	data, err := s.repo.Insert(data)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *reservation_service) Update(id uint, res *models.Reservation) *library.Responses {
	data, err := s.repo.Set(id, res)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *reservation_service) Delete(id uint) *library.Responses {
	data, err := s.repo.Remove(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}