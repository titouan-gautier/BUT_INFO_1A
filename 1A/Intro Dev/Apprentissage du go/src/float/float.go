package main

import "fmt"

func main() {
	var notes []float32 = []float32 {15,16.5,19,13.5,14.5}
	var sommes float32
	for i := 0; i < len(notes); i++ {
		sommes = sommes + notes[i]
	}
	fmt.Println(sommes/5)
	fmt.Println((15+16.5+19+13.5+14.5)/5)
}
