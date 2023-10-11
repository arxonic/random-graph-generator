package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

type Graph struct {
	nodesNum int
	edgesNum int
	matrix   [][]int
}

func (g *Graph) GraphGenerator(avrgConNum int) {
	for i := 0; i < g.nodesNum; i++ {
		relatedNodesNum := PoissonRandom(avrgConNum, g.nodesNum)
		relatedNodes := g.getRandomNodes(relatedNodesNum, i)
		g.matrix[i] = relatedNodes
		g.edgesNum += relatedNodesNum
	}
}

func (g *Graph) getRandomNodes(relatedNodesNum, currentNode int) []int {
	randomNodes := make([]int, g.nodesNum-1)
	for i := 0; i < len(randomNodes); i++ {
		if i >= currentNode {
			randomNodes[i] = i + 1
			continue
		}
		randomNodes[i] = i
	}

	rand.Shuffle(len(randomNodes), func(i, j int) { randomNodes[i], randomNodes[j] = randomNodes[j], randomNodes[i] })

	return randomNodes[:relatedNodesNum]
}

func (g *Graph) PrintGraph() {
	fmt.Println("Graph size: ", g.nodesNum, "\nEdges num: ", g.edgesNum, "\nList view: ")
	for i := 0; i < g.nodesNum; i++ {
		fmt.Println("(", i, ")\t", g.matrix[i])
	}
}

func (g *Graph) CSVWriter(fileName string) {
	s := make([][]string, g.edgesNum+1)
	k := 0
	s[k] = append(s[k], "src")
	s[k] = append(s[k], "dst")
	k++
	for i := 0; i < g.nodesNum; i++ {
		for j := 0; j < len(g.matrix[i]); j++ {
			s[k] = append(s[k], strconv.Itoa(i))
			s[k] = append(s[k], strconv.Itoa(g.matrix[i][j]))
			k++
		}
	}

	f, err := os.Create(fileName + ".csv")
	if err != nil {
		fmt.Println("Problem with file: ", err)
		return
	}

	w := csv.NewWriter(f)
	w.WriteAll(s)

	if err := w.Error(); err != nil {
		fmt.Println("Error writing csv: ", err)
	}
}

func NewGraph(nodesNum int) Graph {
	g := Graph{}
	g.nodesNum = nodesNum
	g.matrix = make([][]int, nodesNum)
	return g
}
