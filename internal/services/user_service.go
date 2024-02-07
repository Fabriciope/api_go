package services

import (
	"errors"
	"time"

	"github.com/Fabriciope/my-api/configs"
	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/models"
)

// TODO: testar
type UserService struct {
	repository repositories.RepositoryInterface
}

func NewUserService(repository repositories.RepositoryInterface) *UserService {
	return &UserService{repository}
}

func (us *UserService) CreateUser(dto *dto.CreateUserInput) error {
	user, err := models.NewUser((*dto).Name, (*dto).Email, (*dto).Password)
	if err != nil {
		return err
	}

	if us.repository.Create(user) != nil {
		return errors.New("Error when inserting user: " + user.Name)
	}

	return nil
}

func (us *UserService) AuthenticateUser(dto dto.GetJWTInput) (token string, err error) {
	u, err := us.repository.FindOneWhere("email", dto.Email)
	if err != nil {
		return "", errors.New("email not registered")
	}

	user := u.(*models.User)
	if !user.ValidatePassword(dto.Password) {
		return "", errors.New("incorrect password")
	}

	jta := configs.Cfg.JWTTokenAuth
	claims := map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Minute * time.Duration(configs.Cfg.JWTExpiresIn)).Unix(),
	}
	_, token, err = jta.Encode(claims)
	if err != nil {
		return "", err
	}

	return
}
