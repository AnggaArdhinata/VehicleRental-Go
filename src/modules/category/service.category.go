package category

import (
	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type category_service struct {
	repo interfaces.Category_RepoIF
}

func NewService(repo interfaces.Category_RepoIF) *category_service {
	return &category_service{repo}
}

func (s *category_service) FindById(id uint) *library.Responses {
	data, err := s.repo.GetById(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	
	return library.Response(data, 200, false)
}

func (s *category_service) FindAll() *library.Responses {
	data, err := s.repo.GetAll()
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	
	return library.Response(data, 200, false)
}

func (s *category_service) Add(data *models.Category) *library.Responses {
	data, err := s.repo.Insert(data)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *category_service) Update(id uint, res *models.Category) *library.Responses {
	data, err := s.repo.Set(id, res)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *category_service) Delete(id uint) *library.Responses {
	data, err := s.repo.Remove(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}