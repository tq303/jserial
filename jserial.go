package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "strconv"
)

type framesArray [][][]string

func main() {

	jsonData, err := readJsonFile("./json/cascade-down.json");
    
    if err != nil {
        log.Fatal(err)
    }

    for i := range jsonData {
    	for j := range jsonData[i] {
    		for k := range jsonData[i][j] {
        		fmt.Println("Output to Serial", strconv.ParseInt(jsonData[i][j][k], 10))
        	}
    	}
    }
}

func readJsonFile(fileName string) ([][][]string, error) {
	
	var jsonData framesArray

    jsonFile, err := ioutil.ReadFile(fileName)

    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(jsonFile, &jsonData)

    if err != nil {
        return nil, err
    }

    return jsonData, nil
}