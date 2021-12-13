package lib

import "fmt"

type Liste struct {
	caps *Element
}

type Element struct {
	v int
	next *Element
	prev *Element
}

func (l *Liste) Head() int {
	return l.caps.v
}

func (l Liste) Tail() Liste {
	return Liste{caps: l.caps.next}
}

func (l *Liste) Append(e int){
	x := Element{v : e, next: l.caps, prev: nil}  
	if l.caps != nil {
		l.caps.prev = &x
	}
	l.caps = &x
}

func ListeVide() Liste {
	return Liste{nil}
}

func (l Liste) VerifVide() bool {
	return l.caps.next == nil
}

func (l Liste) Affiche() string {
	var s string
	for !l.VerifVide() {
		s = fmt.Sprint(s, "", l.Head())
		l = l.Tail()
	}
	return s
}
