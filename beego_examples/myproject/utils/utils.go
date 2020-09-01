package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "chris"
	password = "chris"
	dbname   = "demo"
)

func getPsqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func getDB() *sql.DB {
	db, err := sql.Open("postgres", getPsqlInfo())
	if err != nil {
		log.Fatal(err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("successfull connected!")
	return db
}

var db = getDB()

// InitPostgres =>
func InitPostgres() {
	if db == nil {
		db, _ := sql.Open("postgres", getPsqlInfo())
		fmt.Println(db)
		CreateTableWithUser()
	}
}

// ModifyDB =>
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// CreateTableWithUser =>
func CreateTableWithUser() {
	sql := `create table if not exists beego_users (
			id serial primary key not null,
			username varchar(64) not null,
			password varchar(64) not null,
			status int,
			createtime int
	);`
	fmt.Println(sql)
	ModifyDB(sql)
}

// QueryRowDB =>
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// MD5 =>
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

func init() {
	InitPostgres()
	CreateTableWithUser()
}
