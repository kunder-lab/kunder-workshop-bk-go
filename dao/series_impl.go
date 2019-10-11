package dao

import (
	"kunder-workshop-bk-go/dbConnections"
	"kunder-workshop-bk-go/models"
)

type SerieImpl struct {
}

func (dao SerieImpl) Create(serie models.Example) error {

	query := "INSERT INTO serie (id) VALUES( ?)"

	stmt, err := dbConnections.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(serie.ID)

	if err != nil {
		return err
	}

	return nil
}
