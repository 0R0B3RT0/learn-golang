package main

import (
	"bytes"
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
	c := cachorro{"Padawan", "SRD", 6}
	fmt.Println(c)

	cachorroJson, error := json.Marshal(c)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(cachorroJson)

	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]interface{}{
		"nome":  "Rex",
		"raca":  "Pastor Alemão",
		"idade": 3,
	}

	cachorroJson2, error2 := json.Marshal(c2)
	if error2 != nil {
		log.Fatal(error)
	}

	fmt.Println(cachorroJson2)

	fmt.Println(bytes.NewBuffer(cachorroJson2))
}
