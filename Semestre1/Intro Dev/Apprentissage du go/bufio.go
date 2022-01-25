package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// Ouverture du fichier
	var filePath string = "fichierScan"
	var myFile *os.File
	var err error
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Préparation de la lecture
	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(myFile)

	// Lecture des lignes du fichier
	for scanner.Scan() {
		log.Print("Je viens de lire : ", scanner.Text())
	}

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
