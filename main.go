package main

import (
	"encoding/json"
	"fmt"
)

type ItemDetail struct {
	Data string `json:"data"`
}

type Item struct {
	Name      string      `json:"name"`
	Parameter interface{} `json:"parameter"`
}

func main() {
	data := `{
    "name": "1234",
    "parameter": {
      "data":"1234"
    }
  }`

	cItem := Item{}
	json.Unmarshal([]byte(data), &cItem)

	fmt.Println(cItem)
	detail := ItemDetail{}

	fmt.Println(detail.Data)
}
