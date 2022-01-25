package lib

type Arbre struct {
	valeur int
	tab []*Arbre
}

func (a *Arbre) Root() int {
	return a.valeur
}

func (a *Arbre) Make(v int,b []*Arbre) {
	a.valeur = v
	a.tab = b
}

func ArbreVide() Arbre {
	return Arbre{tab: nil}
}

func (a Arbre) IsEmpty() bool {
	return a.tab == nil
}