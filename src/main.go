package main

import (
	"fmt"

	"github.com/cnhmk/golangMusicDB/src/helloworld"
	"github.com/cnhmk/golangMusicDB/src/jsonGenerator"
)

func main() {
	fmt.Println(helloworld.HelloWorld())
	fmt.Println(jsonGenerator.GenerateJSON())
}
