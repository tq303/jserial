package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {

	jsonFile, err := ioutil.ReadFile("cascade-down.json")

	if err != nil {
		log.Fatal(err)
	}

}

func ListFiles() {

	files, err := ioutil.ReadDir("./json")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}