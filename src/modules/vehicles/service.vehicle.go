package vehicles

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type vehicle_service struct {
	repo interfaces.Vehicles_RepoIF
}

func NewService(repo interfaces.Vehicles_RepoIF) *vehicle_service {
	return &vehicle_service{repo}
}

func (s *vehicle_service) FindAll() *library.Responses {
	data, err := s.repo.GetAll()
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *vehicle_service) Add(data *models.Vehicle) *library.Responses {
	data, err := s.repo.Insert(data)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *vehicle_service) Update(id uint, res *models.Vehicle) *library.Responses {
	data, err := s.repo.Set(id, res)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *vehicle_service) Delete(id uint) *library.Responses {
	data, err := s.repo.Remove(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}