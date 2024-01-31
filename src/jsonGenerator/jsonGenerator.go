package jsonGenerator

import (
	"encoding/json"
	"log"
	"os"
	"strings"
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

func GenerateJSON(
	inputFolder string,
	defaultCategory string,
	description string,
	tag string,
) []byte {
	// Declare required variables:
	var finalResult []DescriptionList
	var unsortedFiles []string
	var unsortedCategory ContentsType

	// User inputs, default values:
	if inputFolder == "" {
		inputFolder = "./"
	}
	if defaultCategory == "" {
		defaultCategory = "Unsorted"
	}
	if description == "" {
		description = "Generated automatically."
	}
	if tag == "" {
		tag = "golang_script"
	} else {
		tag = strings.ReplaceAll(tag, " ", "_")
	}

	// Read files from input folder:
	inputFolderFiles, err := os.ReadDir(inputFolder)
	if err != nil {
		log.Fatal(err)
	}

	// Build a DescriptionList for each folder:
	for _, file := range inputFolderFiles {

		if !file.IsDir() {
			// If it's a file with no folder, append it to other unsorted files:
			unsortedFiles = append(unsortedFiles, file.Name())
		} else {
			// If it's a folder, create a DescriptionList for it:
			currentListName := file.Name()
			var currentUnsortedFiles []string

			// Check each file inside this DescriptionList:
			dlpath := "./" + inputFolder + "/" + currentListName
			currentFolderFiles, err := os.ReadDir(dlpath)
			if err != nil {
				log.Fatal(err)
			}
			var currentDescListContents []ContentsType
			for _, list_file := range currentFolderFiles {

				if !list_file.IsDir() {
					// If it's a file, add it to "unsorted" category of current DescriptionList:
					currentUnsortedFiles = append(currentUnsortedFiles, list_file.Name())
				} else {
					// If it's a folder, create a new ContentsType for this DesctiprionList:
					currentContentsName := list_file.Name()
					var currentContentsEntries []string

					ctpath := "./" + inputFolder + "/" + currentListName + "/" + currentContentsName
					currentContentsFiles, err := os.ReadDir(ctpath)
					if err != nil {
						log.Fatal(err)
					}
					for _, entry_file := range currentContentsFiles {
						// Add entries to the current ContentsType:
						currentContentsEntries = append(currentContentsEntries, entry_file.Name())
					}

					// Finish checking files inside this DescriptionList:
					currentContents := ContentsType{
						Category: currentContentsName,
						Entries:  currentContentsEntries,
					}
					currentDescListContents = append(currentDescListContents, currentContents)
				}
			}
			// Add unsorted files in this DescriptionList to a new ContentType:
			if len(currentUnsortedFiles) != 0 {
				currentUnsorted := ContentsType{
					Category: defaultCategory,
					Entries:  currentUnsortedFiles,
				}
				currentDescListContents = append(currentDescListContents, currentUnsorted)
			}

			// Finish creating this DescriptionList:
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

			// Add this DescriptionList to the final []DescriptionList:
			finalResult = append(finalResult, currentFolder)
		}
	}

	// Create DescriptionList for unsorted files in the parent folder:
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

	// Add unsorted DescriptionList to []DescriptionList:
	finalResult = append(finalResult, unsortedFolder)

	// Return result as JSON
	//myResultMarshal, _ := json.Marshal(finalResult)
	myResultMarshal, _ := json.MarshalIndent(finalResult, "", " ")
	return []byte(myResultMarshal)
}
