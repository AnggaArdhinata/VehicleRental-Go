package auth

import (

	"github.com/AnggaArdhinata/backend-go/src/databases/orm/models"
	"github.com/AnggaArdhinata/backend-go/src/interfaces"
	"github.com/AnggaArdhinata/backend-go/src/library"
)

/**

** Modules Auth tidak membuat file repo, karena data diambil dari repo user,
** oleh karena itu, struct auth_service cukup mengambil interface user repo.
**/

type auth_service struct {
	repo interfaces.Users_RepoIF
}

type token_response struct {
	Token string `json:"token"`
}

func NewService(repo interfaces.Users_RepoIF) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *models.User) *library.Responses {
	user, err := s.repo.FindByEmail(body.Email)
	if err != nil {
		return library.Response("email atau username tidak valid", 401, true)
	}
	
	//** Compare password jika function "CheckPassword" bernilai false maka return password salah !
	if !library.CheckPassword(user.Password, body.Password) {
		return library.Response("password salah", 401, true)
	}
	jwt := library.NewToken(user.Id, user.Role)
	token, err := jwt.Create()
	if err != nil {
		return library.Response("email belum terdaftar", 500, true)
	}
	return library.Response(token_response{Token: token}, 200, false)
	
}

func (s *auth_service) Register(data *models.User) *library.Responses{
	email, err := s.repo.FindByEmail(data.Email)
	if err != nil {
		return library.Response(err.Error(), 400, true)
	}
	if len(email.Email) > 0 {
		return library.Response("email already registered", 400, true)
	}
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
