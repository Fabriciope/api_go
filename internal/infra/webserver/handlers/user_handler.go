package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Fabriciope/my-api/internal/dto"
	"github.com/Fabriciope/my-api/internal/infra/database/repositories"
	"github.com/Fabriciope/my-api/internal/services"
)

type userHandler struct {
	repository repositories.RepositoryInterface
	service    *services.UserService
}

func newUserHandler(repository repositories.RepositoryInterface) *userHandler {
	return &userHandler{
		repository: repository,
		service:    services.NewUserService(repository),
	}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson("invalid parameters"))
		return
	}

	err = uh.service.CreateUser(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(errorToJson(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(successToJson("user created"))
}

func (uh *userHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var JWTDTO dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&JWTDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(dto.DefaultOutput{
			Error: true,
			Message: "invalid parameters",
		}.ToJson())
		return
	}

	token, err := uh.service.AuthenticateUser(JWTDTO)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(dto.DefaultOutput{
			Error: true,
			Message: err.Error(),
		}.ToJson())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(successWithDataToJson(
		"authenticated",
		dto.GetJWTOutput{
			AccessToken: token,
		},
	))
}
