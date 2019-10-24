package main

import "fmt"
import "encoding/json"

var DATA string = ` 
              { "estimate":{
	                "minutes": "2222",
	                "direction": "North"
	              }
              }`

type Estimates struct {
	Minutes   string `json:"minutes"`
	Direction string `json:"direction"`
}

type Etd struct {
	Est Estimates `json:"estimate"`
}

func main() {
	//fmt.Println(DATA)
	var etd Etd
	json.Unmarshal([]byte(DATA), &etd)
	fmt.Printf("%+v\n", etd)
}
