package main

import (
	"io/ioutil"
	"log"
)

func main() {
	var filePath string = "Table de Multiplication"
	var data []byte
	var err error
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("J'ai lu : ", string(data))
}
