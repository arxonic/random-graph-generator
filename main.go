package main

import (
	"fmt"
)

func userInput() (int, int) {
	nodesNum, avrgConNum := 2, 1
	for {
		fmt.Print("Input num of nodes: ")
		fmt.Scan(&nodesNum)

		if nodesNum < 2 {
			fmt.Println("Num of nodes must be > 1")
			continue
		}

		fmt.Print("Input avrg num of connections: ")
		fmt.Scan(&avrgConNum)

		if avrgConNum < 1 {
			fmt.Println("Avrg num of connections must be > 0")
			continue
		}

		break
	}

	return nodesNum, avrgConNum
}

func main() {
	nodesNum, avrgConNum := userInput()

	graph := NewGraph(nodesNum)
	graph.GraphGenerator(avrgConNum)
	graph.PrintGraph()
	graph.CSVWriter("graph")

	graph2 := NewGraph(nodesNum)
	graph2.GraphGenerator(avrgConNum)
	graph2.PrintGraph()
	graph2.CSVWriter("graph2")
}
