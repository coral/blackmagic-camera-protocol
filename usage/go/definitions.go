package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BlackmagicCameraProtocol struct {
	Information struct {
		Readme string `json:"readme"`
		Source string `json:"source"`
		Git    string `json:"git"`
	} `json:"information"`
	Groups            []Group            `json:"groups"`
	BluetoothServices []BluetoothService `json:"bluetooth_services"`
}

type Group struct {
	Name           string      `json:"name"`
	NormalizedName string      `json:"normalized_name"`
	ID             int         `json:"id"`
	Parameters     []Parameter `json:"parameters"`
}

type Parameter struct {
	ID             int      `json:"id"`
	Group          string   `json:"group"`
	GroupID        int      `json:"group_id"`
	Parameter      string   `json:"parameter"`
	Type           string   `json:"type"`
	Index          []string `json:"index"`
	Minimum        float32  `json:"minimum,omitempty"`
	Maximum        float32  `json:"maximum,omitempty"`
	Interpretation string   `json:"interpretation"`
}

type BluetoothService struct {
	Name            string           `json:"name"`
	NormalizedName  string           `json:"normalized_name"`
	UUID            string           `json:"uuid"`
	Characteristics []Characteristic `json:"characteristics"`
}

type Characteristic struct {
	Name           string `json:"name"`
	NormalizedName string `json:"normalized_name"`
	UUID           string `json:"uuid"`
	Description    string `json:"description"`
}

func main() {
	definition, _ := ioutil.ReadFile("../../PROTOCOL.json")
	var decoded BlackmagicCameraProtocol

	err := json.Unmarshal(definition, &decoded)
	if err != nil {
		panic(err)
	}

	fmt.Println(decoded)
}
