package lib

type Files struct {
	premier *Element
	dernier *Element

}

type Element struct {
	val int
	prev *Element
	next *Element
}

func (f *Files) Push() Files {

}

func (f *Files) Pull() int {
	
}