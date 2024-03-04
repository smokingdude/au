package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type RawBom struct {
	Components []struct {
		Type     string `json:"type"`
		BomRef   string `json:"bomRef"`
		Name     string `json:"name"`
		Version  string `json:"version"`
		Scope    string `json:"scope"`
		Purl     string `json:"purl"`
		Hashes   string `json:"hashes"`
		Licences []struct {
			Id string `json:"id"`
		} `json:"licences"`
	} `json:"components"`
}

func main() {
	result := RawBom{}
	path := "./bom.json"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		fmt.Println(err)
	}
	for _, comp := range result.Components {
		fmt.Println(comp.Name, comp.Type, comp.Licences)
	}
}
