package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct{
	db *sqlx.DB
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	return nil,nil
}

func (r accountRepositoryDB) GetAll(int) ([]Account, error) {
	return nil, nil
}