package utils

import (
	"encoding/json"
	"log"
)

func PrettyPrint(x interface{}){
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		log.Println("error:", err)
	}
	log.Println(string(b))
}