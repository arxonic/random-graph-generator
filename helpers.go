package main

import (
	"math"
	"math/rand"
)

func PoissonRandom(avrgConnectiomNum, graphSize int) int {
	l := math.Exp(float64(-avrgConnectiomNum))
	p := 1.0
	k := 0
	for {
		p = p * rand.Float64()
		k++
		if p <= l || k == graphSize-1 {
			break
		}
	}
	if k <= 1 {
		return 1
	}
	return k - 1
}
