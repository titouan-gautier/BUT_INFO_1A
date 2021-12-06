package acrostiche

import (
	"bufio"
	"log"
	"os"
)

/*
La fonction acrostiche doit reconstituer le mot formé par le premier caractère
de chaque ligne d'un fichier, en ignorant les lignes vides.

# Entrée
- fName : le nom d'un fichier

# Sortie
- mot : la chaîne de caractère obtenue en mettant l'une après l'autre, dans
        l'ordre des lignes du fichier, les premières lettres de chacunes des
        lignes du fichier dont le nom est fName (on considère que ce fichier
        existe toujours)
*/
func acrostiche(fName string) (mot string) {
	var MyFile *os.File
	var err error

	MyFile, err = os.Open(fName)

	if err != nil {
		log.Fatal()
	}

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(MyFile)
	for scanner.Scan() {
		var s string = scanner.Text()
		if s != "" {
			mot = mot + string(s[0])
		}
	}

	err = MyFile.Close()
	if err != nil {
		log.Fatal()
	}

	return mot
}
