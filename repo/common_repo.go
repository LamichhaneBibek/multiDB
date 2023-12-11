package repo

import "github.com/LamichhaneBibek/multiDB/database"

type DBOperator struct {
	Con database.DBConnection
}

func (db *DBOperator) Save(value interface{}) error {
	return db.Con.Statement.Create(value).Error
}

func (db *DBOperator) Find(dest interface{}) error {
	return db.Con.Statement.Find(dest).Error
}
