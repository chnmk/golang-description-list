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

	// Declare required variables:
	var finalResult []DescriptionList
	var unsortedFiles []string
	var unsortedCategory ContentsType

	// User inputs, default values:
	inputFolder := "temp_with_folders"
	defaultCategory := "Unsorted"
	// Leave empty to exclude:
	description := "Generated automatically."
	tag := "golang_script"

	// Remove spaces from tag input:
	// ...

	// Read files from the folder:
	inputFolderFiles, err := os.ReadDir("./" + inputFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Build a DescriptionList for each folder:
	for _, file := range inputFolderFiles {

		if !file.IsDir() {
			// = songs with no artists
			// If it's a file with no folder, append it to other unsorted files:
			unsortedFiles = append(unsortedFiles, file.Name())
		} else {
			// = artist
			// If it's a folder, create a DescriptionList for it:
			currentListName := file.Name()
			var currentUnsortedFiles []string

			// Check each folder inside this folder:
			currentFolderFiles, err := os.ReadDir("./" + inputFolder + "/" + currentListName)
			if err != nil {
				log.Fatal(err)
			}
			var currentDescListContents []ContentsType
			for _, list_file := range currentFolderFiles {
				// = album

				if !list_file.IsDir() {
					// = songs with no albums
					currentUnsortedFiles = append(currentUnsortedFiles, list_file.Name())
				} else {
					// = album folders
					currentContentsName := list_file.Name()

					var currentContentsEntries []string
					currentContentsFiles, err := os.ReadDir("./" + inputFolder + "/" + currentListName + "/" + currentContentsName)
					if err != nil {
						log.Fatal(err)
					}
					for _, entry_file := range currentContentsFiles {
						// = songs inside albums
						currentContentsEntries = append(currentContentsEntries, entry_file.Name())
					}

					currentContents := ContentsType{
						Category: currentContentsName,
						Entries:  currentContentsEntries,
					}

					currentDescListContents = append(currentDescListContents, currentContents)
					// = finished songs with albums
				}
			}
			if len(currentUnsortedFiles) != 0 {
				currentDescListUnsortedContents := ContentsType{
					Category: defaultCategory,
					Entries:  currentUnsortedFiles,
				}
				currentDescListContents = append(currentDescListContents, currentDescListUnsortedContents)
			}

			currentFolder := DescriptionList{
				Name: currentListName,
			}
			if description != "" {
				currentFolder.Description = description
			}
			if tag != "" {
				currentFolder.Tags = []string{tag}
			}
			if len(currentDescListContents) != 0 {
				currentFolder.Contents = currentDescListContents
			}
			finalResult = append(finalResult, currentFolder)
		}
	}

	// Add this to final []DescriptionList
	unsortedCategory = ContentsType{
		Category: defaultCategory,
		Entries:  unsortedFiles,
	}
	unsortedFolder := DescriptionList{
		Name: defaultCategory,
	}
	if description != "" {
		unsortedFolder.Description = description
	}
	if tag != "" {
		unsortedFolder.Tags = []string{tag}
	}
	if len(unsortedCategory.Entries) != 0 {
		unsortedFolder.Contents = []ContentsType{unsortedCategory}
	}

	finalResult = append(finalResult, unsortedFolder)

	// Return result as JSON
	myResultMarshal, _ := json.Marshal(finalResult)
	return string(myResultMarshal)
}
