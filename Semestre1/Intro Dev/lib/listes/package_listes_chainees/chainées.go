package lib

import (
	"fmt"
)

type Liste_chainée struct {
	head int
	tail *Liste_chainée
}

func (l Liste_chainée) Head() int {
	return l.head
}

func (l Liste_chainée) Tail() Liste_chainée {
	return *l.tail
}

func Append(e int, l Liste_chainée) Liste_chainée {
	return Liste_chainée{head: e, tail: &l}
}

func ListeVide() Liste_chainée {
	return Liste_chainée{tail: nil}
}

func (l Liste_chainée) VerifVide() bool {
	return l.tail == nil
}

func (l Liste_chainée) Affiche() string {
	var s string
	for !l.VerifVide() {
		s = fmt.Sprint(s, " ", l.Head())
		l = l.Tail()
	}
	return s[1:]
}
