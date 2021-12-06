package lignes

/*
La fonction lignes doit compter le nombre de lignes non vides dans un
fichier dont le nom est indiqué en paramètre.

# Entrée
- fName : une chaîne de caractères correspondant à un nom de fichier

# Sortie
- nLignes : un entier indiquant le nombre de lignes non vides dans le
fichier de nom fName ou -1 si le fichier n'existe pas
*/

import (
	"bufio"
	"os"
)

func lignes(fName string) (nLignes int) {
	var MyFile *os.File
	var err error

	MyFile, err = os.Open(fName)

	if err != nil {
		return -1
	}

	var scanner *bufio.Scanner
	scanner = bufio.NewScanner(MyFile)
	for scanner.Scan() {
		if scanner.Text() != "" {
			nLignes++
		}
	}

	err = MyFile.Close()
	if err != nil {
		return -1
	}
	return nLignes
}
