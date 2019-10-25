package main

import "fmt"
import "encoding/json"

var DATA string = ` 
              {
				"name": "Montgomery St.",
				"abbr": "MONT",
				"etd": [{
					"destination": "Antioch",
					"abbreviation": "ANTC",
					"limited": "0",
					"estimate":[ {
	            	    "minutes": "2222",
	            	    "direction": "North"
	            	  },{
	            	    "minutes": "3333",
	            	    "direction": "South"
						}]
				}]}`

type Estimates []struct {
	Minutes   string `json:"minutes"`
	Direction string `json:"direction"`
}

type Etd []struct {
	Destination  string    `json:"destination"`
	Abbreviation string    `json:"abbreviation"`
	Limited      int       `json:"limited"`
	Est          Estimates `json:"estimate"`
}

type Station struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	Etd  Etd    `json:"etd"`
}

func main() {
	//fmt.Println(DATA)
	//var etd Etd
	var station Station
	json.Unmarshal([]byte(DATA), &station)
	fmt.Printf("%+v\n", station)
	fmt.Println(station)
}
