package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "encoding/hex"
  "net/http"
)

type framesArray [][][]string

type Animation struct {
  Frames framesArray
}

func main() {

    initEndpoints()

}

func convert() {

  jsonData, err := readJsonFile("./json/cascade-down.json")

  if err != nil {
    log.Fatal(err)
  }

  loopOut(jsonData)
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

func initEndpoints() {

  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {

    var body Animation

    err := json.NewDecoder(r.Body).Decode(&body)

    if err != nil {
      log.Println(err)
    }

    loopOut(body.Frames)

    fmt.Fprintf(w, "handler")

  })

  http.ListenAndServe(":8080", nil)
}

func loopOut(jsonData framesArray) {
  for i := range jsonData {
    for j := range jsonData[i] {
      for k := range jsonData[i][j] {

        colour, err := hex.DecodeString(jsonData[i][j][k])

        // reset to black if there is an error
        if err != nil {
          colour = make([]byte, 3)
        }

        fmt.Println("Output to Serial", colour[0], colour[1], colour[2])
      }
    }
  }
}