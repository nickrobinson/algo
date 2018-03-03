package mst

import (
	"fmt"
	"math"

	"github.com/nickrobinson/algo/pq"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

var nodeMap = map[int]string{
	0: "a",
	1: "b",
	2: "c",
	3: "d",
	4: "e",
	5: "f",
	6: "g",
	7: "h",
	8: "i",
}

// Kruskals Algorithm Implementation
func Kruskals(g *simple.WeightedUndirectedGraph) *simple.WeightedUndirectedGraph {
	ds := newDisjointSet()

	queue := pq.New()
	mst := simple.NewWeightedUndirectedGraph(0, math.Inf(1))

	for _, n := range g.Nodes() {
		ds.makeSet(n.ID())
	}

	for _, e := range g.Edges() {
		weight, _ := g.Weight(e.From(), e.To())
		queue.Insert(e, weight)
	}

	for queue.Len() > 0 {
		edge, _ := queue.Pop()
		test := edge.(graph.Edge)

		// If two elements are part of different sets then it is ok to join
		if ds.find(test.From().ID()) != ds.find(test.To().ID()) {
			ds.union(ds.find(test.From().ID()), ds.find(test.To().ID()))

			fmt.Println("Adding Edge: " + nodeMap[int(test.From().ID())] + "->" + nodeMap[int(test.To().ID())])

			weight, _ := g.Weight(test.From(), test.To())

			mst.SetWeightedEdge(simple.WeightedEdge{
				F: simple.Node(test.From().ID()),
				T: simple.Node(test.To().ID()),
				W: weight})
		} else {
			fmt.Println("Not Adding Edge: " + nodeMap[int(test.From().ID())] + "->" + nodeMap[int(test.To().ID())])
		}
	}

	return mst
}
