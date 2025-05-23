package main

import (
    "os"
    "encoding/json"
	"fmt"
)

func loadChoicesFromJSON(filename string) ([]string, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var obj map[string]interface{}
    if err := json.Unmarshal(data, &obj); err != nil {
        return nil, err
    }

    choices := make([]string, 0, len(obj))
    for k := range obj {
        choices = append(choices, k)
    }
    return choices, nil
}

func main(){
	fmt.Println("Hello")
	books, _ :=loadChoicesFromJSON("./../bible/kjv.json")
	fmt.Println(books)
}