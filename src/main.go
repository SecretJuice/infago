package main

import (
	"fmt"
	infago "infago/src/infago"
)

func main() {

	mysql := infago.ConnectMySQL()
	rows, err := mysql.Query("SELECT * FROM classes;")
	if err != nil {
		fmt.Errorf("ERROR: QUERY FAILED")
		panic(err)
	}

	for rows.Next() {
		var id int
		var class_name string
		var student_id int
		if err := rows.Scan(&id, &class_name, &student_id); err != nil {
			panic(err)
		}
		fmt.Printf("Class ID: %d, Class Name: %s, Student ID: %d\n", id, class_name, student_id)
	}
}
