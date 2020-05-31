package utils

import (
	"github.com/go-pg/pg"
)

func Connect_DB(user string, password string, database string) *pg.DB {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: database,
	})
	return db
}
