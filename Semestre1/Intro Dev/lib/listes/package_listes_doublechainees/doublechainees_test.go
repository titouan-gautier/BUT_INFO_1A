package lib

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	var l = ListeVide()
	log.Print(l.Head())
	log.Print(l.Tail())
}
