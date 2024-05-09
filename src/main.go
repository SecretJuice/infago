package main

import (
	"fmt"
	infago "infago/src/infago"
)

func main() {

	query := `
	SELECT id from classes;
	`

	mysql := infago.ConnectMySQL()
	rows, err := mysql.Query(query)
	if err != nil {
		fmt.Errorf("ERROR: QUERY FAILED")
		panic(err)
	}

	var tables string
	for rows.Next() {
		if err := rows.Scan(&tables); err != nil {
			panic(err)
		}
		fmt.Printf("Table: %s", tables)
	}
}
