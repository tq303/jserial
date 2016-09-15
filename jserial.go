package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
)

type nestedArray [][][]string

func main() {

    var jsonData nestedArray

    jsonFile, err := ioutil.ReadFile("./json/cascade-down.json")

    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(jsonFile, &jsonData)

    if err != nil {
        log.Fatal(err)
    }

    for i := range jsonData {
    	for j := range jsonData[i] {
    		for k := range jsonData[i][j] {
        		fmt.Println("Output to Serial", jsonData[i][j][k])
        	}
    	}
    }
}

func frame() []uint8 {

	strips := 8
	leds := 30
	totalLeds := strips * leds

	return make([]uint8, totalLeds);
}