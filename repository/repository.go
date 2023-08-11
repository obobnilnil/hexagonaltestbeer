package repository

import (
	"app/model"
	"database/sql"
	"log"
)

type IRepository interface {
	Get(req model.GetRequest) ([]model.Data, error)
	Add(req model.AddRequest) (error, *int64)
	Update(query string, args ...interface{}) error
	Delete(req model.DeleteRequest) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Get(req model.GetRequest) ([]model.Data, error) {
	//
	//SQL logic
	//
	data := []model.Data{}
	rows, err := r.db.Query("SELECT id, name, type, detail, url FROM beer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var u model.Data
		err := rows.Scan(&u.ID, &u.Name, &u.Type, &u.Detail, &u.Url)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func (r *Repository) Add(req model.AddRequest) (error, *int64) {
	query := "INSERT INTO beer (name, type, detail, url) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, req.Name, req.Type, req.Detail, req.Url)
	if err != nil {
		return err, nil
	}

	lastID, _ := result.LastInsertId()

	rowCount, err := result.RowsAffected() /// logic for check replicate or not
	if rowCount == 0 {
		return err, nil
	}
	return nil, &lastID
}

func (r *Repository) Update(query string, args ...interface{}) error {
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(req model.DeleteRequest) error {
	query := "DELETE FROM beer WHERE id = ?"
	log.Println(req.ID)
	_, err := r.db.Exec(query, req.ID)
	if err != nil {
		return err
	}

	return nil
}
