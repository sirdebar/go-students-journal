package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func deferFunc() {
	f, err := os.Open(names.txt)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scan.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(data)
	}

}
