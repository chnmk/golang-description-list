package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cnhmk/golangMusicDB/src/jsonGenerator"
)

func main() {

	var inputFolder string
	fmt.Println("Select a folder.")
	fmt.Println("Enter './' to scan this folder,")
	fmt.Println("Or enter a path, such as './my_folder/':")
	fmt.Scan(&inputFolder)

	fmt.Println("Enter default name for a list of unsorted files")
	fmt.Println("(Default: 'Unsorted'):")
	reader := bufio.NewReader(os.Stdin)
	defaultCategory, _ := reader.ReadString('\n')
	defaultCategory = strings.TrimSuffix(defaultCategory, "\n")

	fmt.Println("Enter description for generated lists")
	fmt.Println("(Default: 'Generated automatically.'):")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSuffix(description, "\n")

	fmt.Println("Enter a tag for generated lists")
	fmt.Println("(Default: 'golang_script'):")
	tag, _ := reader.ReadString('\n')
	tag = strings.TrimSuffix(tag, "\n")

	fmt.Println(jsonGenerator.GenerateJSON(inputFolder, defaultCategory, description, tag))
}
