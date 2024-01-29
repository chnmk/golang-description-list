package jsonGenerator

import (
	"encoding/json"
	"log"
	"os"
)

type ContentsType struct {
	Category string   `json:"Category"`
	Entries  []string `json:"Entries,omitempty"`
}

type DescriptionList struct {
	Name        string         `json:"Name"`
	Description string         `json:"Description,omitempty"`
	Tags        []string       `json:"Tags,omitempty"`
	Contents    []ContentsType `json:"Contents,omitempty"`
}

func GenerateJSON() string {

	// Read files from ./temp
	files, err := os.ReadDir("./temp")
	if err != nil {
		log.Fatal(err)
	}

	// Declare required variables
	var filesOutput []string
	var newContent ContentsType

	// Get variables from files
	for _, file := range files {
		filesOutput = append(filesOutput, file.Name())
		newContent = ContentsType{
			Category: "MyCat",
			Entries:  filesOutput,
		}
	}

	// Build a DescriptionList
	myResult := &DescriptionList{
		Name:        "MyName",
		Description: "MyDesc",
		Tags:        []string{"Tag1", "Tag2"},

		Contents: []ContentsType{newContent},
	}

	// Return result
	myResultMarshal, _ := json.Marshal(myResult)
	return string(myResultMarshal)
}
