package main

import (
	"encoding/csv"
	"fmt"
	infago "infago/src/infago"
	"os"
	"strconv"
)

type StudentClass struct {
	StudentId int
	Name      string
	ClassName string
	ClassId   int
}

type ClassByStudent struct {
	ClassId   int
	ClassName string
	StudentId int
}

type Student struct {
	Id       int
	Name     string
	FavColor string
	Age      string
}

func main() {
	students := infago.GetPostgresData()
	classes := infago.GetMySQLData()

	var studentArray []Student
	var student Student
	for students.Next() {
		students.Scan(&student.Id, &student.Name, &student.FavColor, &student.Age)
		studentArray = append(studentArray, student)
	}

	var classArray []ClassByStudent
	var classByStudent ClassByStudent
	for classes.Next() {
		classes.Scan(&classByStudent.ClassId, &classByStudent.ClassName, &classByStudent.StudentId)
		classArray = append(classArray, classByStudent)
	}

	var StudentClasses []StudentClass
	var studentClass StudentClass
	for _, _student := range studentArray {
		for _, _class := range classArray {
			if _student.Id == _class.StudentId {
				studentClass.StudentId = _student.Id
				studentClass.Name = _student.Name
				studentClass.ClassName = _class.ClassName
				studentClass.ClassId = _class.ClassId

				StudentClasses = append(StudentClasses, studentClass)
			}
		}
	}

	for _, value := range StudentClasses {
		fmt.Printf("Class Entry: %d, %s, %d, %s\n", value.ClassId, value.ClassName, value.StudentId, value.Name)
	}

	writeToCSV(StudentClasses)

}

func writeToCSV(studentClasses []StudentClass) {
	file, err := os.Create("./result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	var rows [][]string

	for _, element := range studentClasses {

		row := []string{strconv.Itoa(element.StudentId), element.Name, element.ClassName, strconv.Itoa(element.ClassId)}
		rows = append(rows, row)
	}

	header := []string{"StudentId", "Name", "ClassName", "ClassId"}

	writer.Write(header)
	writer.WriteAll(rows)

}
