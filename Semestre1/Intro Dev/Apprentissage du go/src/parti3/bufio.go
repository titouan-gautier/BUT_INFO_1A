package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
	
)

func main() {
	// Ouverture du fichier
	var filePath string = "notes.csv"
	var myFile *os.File
	var err error
	var s []string
	//var name string
	var note int
	var bonus int
	var tot int
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)

	// Lecture des lignes du fichier
	for i:= 0 ; scanner.Scan() ; i++ {
		s = strings.Split(scanner.Text(),",")
		strings.Split(scanner.Text(),"_")
		//name = s[0]
		note,_ = strconv.Atoi(s[1])
		bonus,_ = strconv.Atoi(s[2])
		tot,_ = strconv.Atoi(s[3])
		if (note+bonus) != tot {
			log.Print("La ligne ", i ," n'est pas bonne")
		}
	}
	log.Print("C'est fini")

	// Vérification que tout s'est bien passé
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
