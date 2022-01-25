package main

import (
	"fmt"
	"log"
	"os"
	"flag"
)

func main() {
	var myFile *os.File
	var err error
	myFile, err = os.Create("Table de Multiplication")
	if err != nil {
		log.Fatal(err)
	}

	var n int 
	flag.IntVar(&n,"n",0,"...")
	flag.Parse()
	if n < 0 || n > 10 {
		fmt.Fprintln(myFile,"ERROR")
	} else {
		for i := 0; i <= 10; i++ {
			_, err = fmt.Fprintln(myFile, n*i)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = myFile.Close()
	if err != nil {
		log.Fatal(err)
	}
}
