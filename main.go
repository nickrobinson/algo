package main

import (
	"fmt"
	"gonum.org/v1/gonum/graph/simple"
	"github.com/nickrobinson/algo/graph/mst"
	"math"
	"strconv"
)

func main() {
	nodes := [14]struct{ srcID, targetID, weight int }{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{7, 8, 7},
		{7, 6, 1},
		{8, 2, 2},
		{8, 6, 6},
		{2, 3, 7},
		{6, 5, 2},
		{2, 5, 4},
		{3, 4, 9},
		{3, 5, 14},
		{5, 4, 10},
	}

	g := simple.NewWeightedUndirectedGraph(0, math.Inf(1))

	for _, n := range nodes {
		g.SetWeightedEdge(simple.WeightedEdge{F: simple.Node(n.srcID), T: simple.Node(n.targetID), W: float64(n.weight)})
	}

	mst := mst.Kruskals(g)
	
	for _, e := range mst.Edges() {
		weight, _ := mst.Weight(e.From(), e.To())
		fmt.Println("From: " + strconv.Itoa(int(e.From().ID())) + 
			" To: " + strconv.Itoa(int(e.To().ID())) +
			" Weight: " + strconv.FormatFloat(weight, 'E', 0, 64))
	}
}