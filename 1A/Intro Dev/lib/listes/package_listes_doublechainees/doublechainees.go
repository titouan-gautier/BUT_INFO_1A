package lib

import "fmt"

type Capsule struct {
	contenu Liste_doublechainée
}

type Liste_doublechainée struct {
	head int
	next *Liste_doublechainée
	prev *Liste_doublechainée
}

func (l Liste_doublechainée) Head() int {
	return l.head
}

func (l Liste_doublechainée) Tail() Liste_doublechainée {
	return *l.next
}

func (l Liste_doublechainée) Prev() Liste_doublechainée {
	return *l.prev
}

func Append(e int, l Liste_doublechainée) Liste_doublechainée {
	return Liste_doublechainée{head: e, next: &l, prev: nil}
}

func ListeVide() Liste_doublechainée {
	return Liste_doublechainée{prev: nil}
}

func (l Liste_doublechainée) VerifVide() bool {
	return l.next == nil
}

func (l Liste_doublechainée) Affiche() string {
	var s string
	for !l.VerifVide() {
		s = fmt.Sprint(s, "", l.Head())
		l = l.Tail()
	}
	return s
}
