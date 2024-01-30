package jsonGenerator

import (
	"encoding/json"
	"fmt"
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

	// Declare required variables:
	var folderName string
	var unsortedFiles []string
	var unsortedCategory ContentsType

	// User inputs:
	inputFolder := "temp"
	defaultCategory := "Various"

	// Read files from the folder:
	files, err := os.ReadDir("./" + inputFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Get variables from each file:
	for _, file := range files {

		if !file.IsDir() {
			unsortedFiles = append(unsortedFiles, file.Name())
			unsortedCategory = ContentsType{
				Category: defaultCategory,
				Entries:  unsortedFiles,
			}
		} else {
			fmt.Println("Hi")
			// For each folder create new ContentsType variable...
		}
	}

	// Build a DescriptionList
	myResult := &DescriptionList{
		Name:        folderName,
		Description: "MyDesc",
		Tags:        []string{"Tag1", "Tag2"},

		Contents: []ContentsType{unsortedCategory},
	}

	// Return result
	myResultMarshal, _ := json.Marshal(myResult)
	return string(myResultMarshal)
}
