package main

import (
	"fmt"
	"cs50_hard/graphs"
)

func main() {

	gr := graphs.New()

	gr.AddNode("0")
	gr.AddNode("B")
	gr.AddNode("A")
	gr.AddNode("1")

	gr.AddLink("0", "A", 6)
	gr.AddLink("0", "B", 2)
	gr.AddLink("B", "A", 3)
	gr.AddLink("A", "1", 1)
	gr.AddLink("B", "1", 5)

	fmt.Println(gr)
	fmt.Println(gr.Path("0", "1"))
}
