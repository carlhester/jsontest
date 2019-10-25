package main

import "fmt"
import "encoding/json"

var DATA string = `
{
  "?xml": {
    "@version": "1.0",
    "@encoding": "utf-8"
  },
  "root": {
    "@id": "1",
    "uri": {
      "#cdata-section": "http://api.bart.gov/api/etd.aspx?cmd=etd&orig=MONT&dir=n&json=y"
    },
    "date": "10/24/2019",
    "time": "11:56:31 AM PDT",
    "station": [
      {
        "name": "Montgomery St.",
        "abbr": "MONT",
        "etd": [
          {
            "destination": "Antioch",
            "abbreviation": "ANTC",
            "limited": "0",
            "estimate": [
              {
                "minutes": "2",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "YELLOW",
                "hexcolor": "#ffff33",
                "bikeflag": "1",
                "delay": "89",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "16",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "YELLOW",
                "hexcolor": "#ffff33",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "31",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "YELLOW",
                "hexcolor": "#ffff33",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              }
            ]
          },
          {
            "destination": "Dublin/Pleasanton",
            "abbreviation": "DUBL",
            "limited": "0",
            "estimate": [
              {
                "minutes": "Leaving",
                "platform": "2",
                "direction": "North",
                "length": "9",
                "color": "BLUE",
                "hexcolor": "#0099cc",
                "bikeflag": "1",
                "delay": "136",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "13",
                "platform": "2",
                "direction": "North",
                "length": "9",
                "color": "BLUE",
                "hexcolor": "#0099cc",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "28",
                "platform": "2",
                "direction": "North",
                "length": "5",
                "color": "BLUE",
                "hexcolor": "#0099cc",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              }
            ]
          },
          {
            "destination": "Richmond",
            "abbreviation": "RICH",
            "limited": "0",
            "estimate": [
              {
                "minutes": "8",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "RED",
                "hexcolor": "#ff0000",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "23",
                "platform": "2",
                "direction": "North",
                "length": "9",
                "color": "RED",
                "hexcolor": "#ff0000",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "38",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "RED",
                "hexcolor": "#ff0000",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              }
            ]
          },
          {
            "destination": "Warm Springs",
            "abbreviation": "WARM",
            "limited": "0",
            "estimate": [
              {
                "minutes": "4",
                "platform": "2",
                "direction": "North",
                "length": "9",
                "color": "GREEN",
                "hexcolor": "#339933",
                "bikeflag": "1",
                "delay": "62",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "18",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "GREEN",
                "hexcolor": "#339933",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              },
              {
                "minutes": "33",
                "platform": "2",
                "direction": "North",
                "length": "10",
                "color": "GREEN",
                "hexcolor": "#339933",
                "bikeflag": "1",
                "delay": "0",
                "carflag": "0",
                "cancelflag": "0",
                "dynamicflag": "0"
              }
            ]
          }
        ]
      }
    ],
    "message": ""
  }
}
`

type Estimates []struct {
	Minutes     string `json:"minutes"`
	Direction   string `json:"direction"`
	Length      int    `json:"length"`
	Color       string `json:"color"`
	Hexcolor    string `json:"hexcolor"`
	Bikeflag    int    `json:"bikeflag"`
	Delay       int    `json:"delay"`
	Carflag     int    `json:"carflag"`
	Cancelflag  int    `json:"cancelflag"`
	Dynamicflag int    `json:"dynamicflag"`
}

type Etd []struct {
	Destination  string    `json:"destination"`
	Abbreviation string    `json:"abbreviation"`
	Limited      int       `json:"limited"`
	Est          Estimates `json:"estimate"`
}

type Station []struct {
	Name string `json:"name"`
	Abbr string `json:"abbr"`
	Etd  Etd    `json:"etd"`
}

type Uri struct {
	Cdata string `json:"#cdata-section"`
}

type Root struct {
	Id      int     `json:"@id"`
	Uri     Uri     `json:"uri"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
	Station Station `json:"station"`
	Message string  `json:"message"`
}

type Xml struct {
	Version  string `json:"@version"`
	Encoding string `json:"@encoding"`
}

type Data struct {
	Xml  Xml  `json:"?xml"`
	Root Root `json:"root"`
}

func main() {
	//fmt.Println(DATA)
	//var etd Etd
	var data Data
	json.Unmarshal([]byte(DATA), &data)
	val := data.Root.Station[0].Etd
	for _, element := range val {
		fmt.Printf("%+v\n\n", element)
	}
	//fmt.Println(val)
	//fmt.Printf("%+v\n", data.Root.Station[0].Name)
	//fmt.Println(data)
}
