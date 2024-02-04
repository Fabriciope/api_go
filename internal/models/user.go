package models

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `json:"id"         db:"id"`
	Name      string    `json:"name"       db:"name"`
	Email     string    `json:"email"      db:"email"`
	Password  string    `json:"-"          db:"password"`
	CreatedAt string    `json:"created_at" db:"created_at"`
}

func NewUser(name, email, password string) (*User, error) {
	// TODO: criptografar a senha quando salvar no banco
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hash),
		CreatedAt: time.Now().String()[:19],
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Get(field string) reflect.Value {
	return reflect.ValueOf(*u).FieldByName(field)
}

func (u *User) DataForDB() map[string]interface{} {
	data := make(map[string]interface{}, 5)
	tUser := reflect.TypeOf(*u)
	fields := reflect.VisibleFields(tUser)
	for _, field := range fields {
		if sField, ok := tUser.FieldByName(field.Name); ok {
			fieldNameInDatabase := string(sField.Tag.Get("db"))
			data[fieldNameInDatabase] = u.Get(field.Name).Interface()
		}
	}
	
	return data
}

func (u User) ToJson() []byte {
	responseJson, err := json.Marshal(u)
	if err != nil {
		return []byte(`{"error": true, "message": "internal server error"}`)
	}

	return responseJson
}
