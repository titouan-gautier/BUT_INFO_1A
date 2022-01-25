package lib

type Piles struct {
	compteur *Element
	sommet *Element
}

type Element struct {
	v int
	next *Element
	prev *Element
}

func (p Piles) Push() Piles {

}

func (p Piles) Pop() Piles {
	Element{}
}