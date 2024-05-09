package infago

import (
	"fmt"
	_"github.com/lib/pq"
	"database/sql"
)

func GetPostgresData() *sql.Rows {
	connStr := "postgresql://pg:password@localhost:5432/students?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Errorf("ERROR: CONNECT DATABASE FAILED")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Errorf("ERROR: PING DATABASE FAILED")
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		fmt.Errorf("ERROR: QUERY FAILED")
		panic(err)
	}

	return rows
}
