package main

import (
  cr821/db-connection
	"fmt"
	infago "infago/src/infago"

)

func main() {

  cr821/db-connection
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
