package main

import (
	"fmt"
	"log"
)

type Student struct {
	Name     string
	Surname  string
	Password string
	Grades   []int
}

type Teacher struct {
	Login    string
	Password string
	Students []Students
}

func main() {
	student := []Student{
		{Name: "Jeremy", Surname: "Scott", Password: "123", Grades: []int{}},
		{Name: "David", Surname: "Lin", Password: "321", Grades: []int{}},
	}
	teacher := []Teacher{
		{Login: "Delayla", Password: "T123", Students: []Students{}},
		{Login: "Margaret", Password: "T123", Students: []Students{}},
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

			if foundTeacher == nil {
				log.Fatal("Teacher not found")
			}

			if foundTeacher.Password != password {
				log.Fatal("Wrong password")
			}

			fmt.Printf("Welcome back, %s", foundTeacher.Login)
		}

	}

}
