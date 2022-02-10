package db

import (
	repo "Main/Repositories"
	"database/sql"
	// "fmt"
	// "github.com/go-pg/pg/v10"
	// "github.com/go-pg/pg/v10/orm"
)

type StockRepo struct {
	db *sql.DB
}

func retrieveAll() {
	// db := pg.Connect()
}

func NewStockRepo(db *sql.DB) *StockRepo {
	return &StockRepo{
		db: db,
	}
}

func (r *StockRepo) FindByID(ID int) (*repo.Data, error) {
	return &repo.Data{}, nil
}
