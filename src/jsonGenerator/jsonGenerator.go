package jsonGenerator

import (
	"encoding/json"
	"log"
	"os"
)

type MyJSON struct {
	ArtistName string `json:"artist"`
	SongName   string `json:"name"`
	FileName   string `json:"file"`
}

func GenerateJSON() string {

	fileNames, err := os.ReadDir("./temp")
	if err != nil {
		log.Fatal(err)
	}

	var fileName []string
	for _, file := range fileNames {
		fileName = append(fileName, file.Name())
	}

	myResult := &MyJSON{
		ArtistName: "My Favourite Artist",
		SongName:   "Best Song Ever",
		FileName:   fileName[1],
	}

	myResultB, _ := json.Marshal(myResult)

	return string(myResultB)
}
