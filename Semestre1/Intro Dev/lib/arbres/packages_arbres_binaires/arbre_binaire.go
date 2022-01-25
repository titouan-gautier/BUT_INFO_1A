package lib

type Arbre struct {
	val int
	gauche *Arbre
	droite *Arbre
}

func (A *Arbre) Root() int {
	return A.val
}

func (A *Arbre) Left() *Arbre {
	return A.gauche
}

func (A *Arbre) Right() *Arbre {
	return A.droite
}

func (A *Arbre) Make(v int,t1 *Arbre, t2 *Arbre) {
	A.val = v
	A.gauche = t1
	A.droite = t2
}

func ArbreVide() Arbre {
	return Arbre{gauche:nil, droite:nil}
}

func (A Arbre) IsEmpty() bool {
	return A.gauche == nil && A.droite == nil
}