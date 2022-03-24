package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

func main() {
	// Ouverture du fichier
	var filePath string = "monfichier"
	var myFile *os.File
	var err error
	myFile, err = os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Lecture d'une ligne
	var nbLus int
	var uneChaine string
	var i int
	nbLus, err = fmt.Fscanln(myFile, &uneChaine)
	for err != io.EOF  {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print("uneChaine = ", uneChaine)
			log.Print("J'ai lu ", nbLus, " valeurs.")
			i = i + nbLus
		}
	nbLus, err = fmt.Fscanln(myFile, &uneChaine)
	}
	log.Print(i)
	// Fermeture du fichier
	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
