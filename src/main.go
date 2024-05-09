package main

import (
	"fmt"
	infago "infago/src/infago"

)

func main() {
	students := infago.GetPostgresData()
	classes := infago.GetMySQLData()

	type Student struct {
		Id int
		Name string
		FavColor string
		Age string
	}

	var studentArray []Student
	var student Student
	for students.Next() {
		students.Scan(&student.Id, &student.Name, &student.FavColor, &student.Age)
		studentArray = append(studentArray, student)
	}

	type ClassByStudent struct {
		ClassId int
		ClassName string
		StudentId int
	}

	var classArray []ClassByStudent
	var classByStudent ClassByStudent
	for classes.Next() {
		classes.Scan(&classByStudent.ClassId, &classByStudent.ClassName, &classByStudent.StudentId)
		classArray = append(classArray, classByStudent)
	}

	type StudentClass struct {
		StudentId int
		Name string
		ClassName string
		ClassId int
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

}
