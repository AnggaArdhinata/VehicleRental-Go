package users

import (

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

type user_service struct {
	repo interfaces.Users_RepoIF
}

func NewService(repo interfaces.Users_RepoIF) *user_service {
	return &user_service{repo}
}

func (s *user_service) FindByImageID(uid string) *library.Responses {
	data, err := s.repo.GetByImgID(uid)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)

}

func (s *user_service) FindById(id uint) *library.Responses {
	data, err := s.repo.GetById(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)

}

func (s *user_service) FindByUserEmail(email string) *library.Responses {
	data, err := s.repo.FindByEmail(email)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)

}

func (s *user_service) FindByUserName(username string) *library.Responses {
	data, err := s.repo.FindByName(username)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)

}

func (s *user_service) FindAll() *library.Responses {
	data, err := s.repo.GetAll()
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *user_service) Add(data *models.User) *library.Responses {
	
	HashPass, err := library.HashPassword(data.Password)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	data.Password = HashPass
	data, err = s.repo.Insert(data)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *user_service) Update(id uint, res *models.User) *library.Responses {
	data, err := s.repo.Set(id, res)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}

func (s *user_service) Delete(id uint) *library.Responses {
	user, err := s.repo.GetById(id)
	if err != nil {
		return library.Response("data not found !", 404, true)
	}
	uidImage := user.ImgId
	library.DeleteImage(uidImage)
	data, err := s.repo.Remove(id)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	return library.Response(data, 200, false)
}