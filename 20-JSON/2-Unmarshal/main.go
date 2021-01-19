package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJson := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	if error := json.Unmarshal([]byte(cachorroJson), &c); error != nil {
		log.Fatal(error)
	}

	fmt.Println(c)

	cachorroJson2 := `{"nome":"Totó","raca":"SND","idade":6}`

	c2 := make(map[string]interface{})
	if error := json.Unmarshal([]byte(cachorroJson2), &c2); error != nil {
		log.Fatal(error)
	}

	fmt.Println(c2)
}
