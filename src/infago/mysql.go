package infago

import (
	"fmt"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySQLData() *sql.Rows {

	db, err := sql.Open("mysql", "root:password@/classes")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query := `
	SELECT * from classes;
	`

	mysql := ConnectMySQL()
	rows, err := mysql.Query(query)
	if err != nil {
		fmt.Errorf("ERROR: QUERY FAILED")
		panic(err)
	}

	return rows
}

func ConnectMySQL() *sql.DB {

	db, err := sql.Open("mysql", "root:password@/classes")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}