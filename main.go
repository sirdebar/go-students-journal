package main

import (
	"fmt"
	"log"
)

type Students struct {
	Name     string
	Surname  string
	Password string
	Grades   []int
}

type Teachers struct {
	Login    string
	Password string
	Students []Students
}

func main() {
	fmt.Print("How do you want to log in?\n1. As Teacher\t 2. As Student\t 3. Exit\n")
	var choice string
	fmt.Scan(&choice)

	switch choice {
	case "3":
		log.Fatal("Exiting...")
	case "1":

	}

}
