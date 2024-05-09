package infago


import (
	"fmt"
	_"github.com/lib/pq"
	"database/sql"
)

func ConnectToPostgres() (*sql.Rows, error) {
	connStr := "postgresql://pg:password@localhost:5432/students?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	return rows, nil
}


// package main
//
// import (
//     "fmt"
// 	"infago/src/infago"
// )
//
// func main() {
//     rows, err := infago.ConnectToPostgres()
//
//     if err != nil {
//         fmt.Printf("Error getting rows: %s\n", err)
//         panic(err)
//     }
//
//
//     var id int
//     var name string
//     var favColor string
//     var age int
//
//     for rows.Next() {
//         rows.Scan(&id, &name, &favColor, &age)
//         fmt.Printf("Id: %d,\n Name: %s,\n Favorite Color: %s,\n Age: %d\n", id, name, favColor, age)
//     }
//
// //     fmt.Printf("Rows: %s", rows)
// }
