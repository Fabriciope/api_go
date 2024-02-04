package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
	defaultActions defaultActions
}

func newUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{
		db: conn,
		defaultActions: defaultActions{
			db: conn,
			table: "users",
		},
	}
}

func (ur *userRepository) Create(user models.ModelInterface) error {
	if !ur.validateModel(user) {
		return errInvalidModel
	}

	err := ur.defaultActions.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(user models.ModelInterface) error {
	if !ur.validateModel(user) {
		return errInvalidModel
	}

	err := ur.defaultActions.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Delete(id uuid.UUID) error {
	return ur.defaultActions.destroy(id)
}

func (ur *userRepository) FindAllWithPagination(page, limit int, sort string) ([]models.ModelInterface, error) {
	sort = strings.ToUpper(sort)
	if sort != "ASC" && sort != "DESC" {
		sort = "ASC"
	} 

	if page <= 0 || limit <= 0 {
		return nil, errPageOrLimitAreWrong
	}

	rows, err := ur.db.Query(fmt.Sprintf(
		"SELECT * FROM users ORDER BY created_at %s LIMIT %d OFFSET %d",
		sort, limit, (page - 1) * limit,
	))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []models.ModelInterface
	for rows.Next() {
		var u models.User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}

func (ur *userRepository) FindOneWhere(attr string, value interface{}) (models.ModelInterface, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE %s = ?", attr)
	stmt, err := ur.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	

	var user models.User
	err = stmt.QueryRow(value).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) validateModel(model models.ModelInterface) bool {
	_, ok := model.(*models.User)
	return ok
}