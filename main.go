package main

import (
	"fmt"
	"log"
)

type Student struct {
	Login    string
	Name     string
	Surname  string
	Password string
	Grades   []int
}

type Teacher struct {
	Login    string
	Password string
	Students []Student
}

func main() {
	student := []Student{
		{Login: "Jeremy1008", Name: "Jeremy", Surname: "Scott", Password: "123", Grades: []int{}},
		{Login: "CoolDavid", Name: "David", Surname: "Lin", Password: "321", Grades: []int{}},
	}
	teacher := []Teacher{
		{Login: "Delayla", Password: "T123", Students: []Student{}},
		{Login: "Margaret", Password: "T123", Students: []Student{}},
	}
	fmt.Print("How do you want to log in?\n1. As Teacher\t 2. As Student\t 3. Exit\n")
	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "3":
		log.Fatal("Exiting...")
	case "1":
		fmt.Print("Type your Login:\n")
		var login string
		fmt.Scan(&login)

		fmt.Print("Type your Password:\n")
		var password string
		fmt.Scan(&password)

		var foundTeacher *Teacher
		for i := range teacher {
			if teacher[i].Login == login {
				foundTeacher = &teacher[i]
				break
			}
		}

		// Проверяем, найден ли учитель (после цикла!)
		if foundTeacher == nil {
			log.Fatal("Teacher not found")
		}

		// Проверяем пароль
		if foundTeacher.Password != password {
			log.Fatal("Wrong password")
		}

		fmt.Printf("Welcome back, %s", foundTeacher.Login)
		// Добавить меню

	case "2":
		fmt.Println("Type your Login:")
		var slogin string
		fmt.Scan(&slogin)

		fmt.Print("Type your Password:\n")
		var spassword string
		fmt.Scan(&spassword)

		var foundStudent *Student
		for i := range student {
			if student[i].Login == slogin {
				foundStudent = &student[i]
				break
			}
		}

		if foundStudent == nil {
			log.Fatal("Student not found") // Исправлено!
		}

		if foundStudent.Password != spassword {
			log.Fatal("Wrong password")
		}

		fmt.Printf("Welcome back, %s %s", foundStudent.Name, foundStudent.Surname)
		// Добавить меню

	}

}
