package repositories

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Fabriciope/my-api/internal/models"
	"github.com/google/uuid"
)

type defaultActions struct {
	db *sql.DB
	table string
}

func (da *defaultActions) Insert(model models.ModelInterface) error {
	attributes, replace, values := getDataToInsert(model)
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		strings.ToLower(da.table), attributes, replace,
	)

	stmt, err := da.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}


	return nil
}

func (da *defaultActions) Update(model models.ModelInterface) error {
	attributesAndReplace, values := getDataToUpdate(model)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", da.table, attributesAndReplace)
	stmt, err := da.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	_, err = stmt.Exec(append(values, model.Get("ID").Interface())...)
	if err != nil {
		return err
	}

	return nil
}

func (da *defaultActions) destroy(id uuid.UUID) error {
	stmt, err := da.db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id = ?", da.table))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func getDataToInsert(model models.ModelInterface) (attributes, replace string, values []interface{}) {
	for field, value := range model.DataForDB() {
		attributes += strings.ToLower(field) + ", "
		replace += "?, "
		values = append(values, value)
	}

	attributes = strings.TrimSuffix(attributes, ", ")
	replace = strings.TrimSuffix(replace, ", ")

	return
}

func getDataToUpdate(model models.ModelInterface) (attributesAndReplace string, values []interface{}) {
	for field, value := range model.DataForDB() {
		if field == "id" {// TODO: criar um slice com os campos que nao devem ser alterador e pular no if
			continue
		}
		attributesAndReplace += field + " = ?, "
		values = append(values, value)
	}

	attributesAndReplace = strings.TrimSuffix(attributesAndReplace, ", ")

	return
}