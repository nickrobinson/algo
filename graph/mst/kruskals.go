package mst

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"github.com/nickrobinson/algo/pq"
	"math"
	"strconv"
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

func Kruskals(g *simple.WeightedUndirectedGraph) (*simple.WeightedUndirectedGraph) {
	set := make(map[int64]int64)
	queue := pq.New()
	mst := simple.NewWeightedUndirectedGraph(0, math.Inf(1))

	for _, n := range g.Nodes() {
		set[n.ID()] = -1
	}

	for _, e := range g.Edges() {
		weight, _ := g.Weight(e.From(), e.To())
		queue.Insert(e, weight)
	}

	for queue.Len() > 0 {
		edge, _ := queue.Pop()
		test := edge.(graph.Edge)

		if(set[test.From().ID()] == -1 && set[test.To().ID()] == -1) {
			set[test.From().ID()] = test.From().ID()
		}

		if(set[test.From().ID()] != set[test.To().ID()]) {
			fmt.Println(nodeMap[int(test.From().ID())] + ":" + strconv.Itoa(int(set[test.From().ID()])) + " " + nodeMap[int(test.To().ID())] + ":" + strconv.Itoa(int(set[test.To().ID()])))

			if(set[test.From().ID()] == -1 && set[test.To().ID()] > 1) {
				set[test.From().ID()] = set[test.To().ID()]
			} else if(set[test.To().ID()] == -1 && set[test.From().ID()] > 1) {
				set[test.To().ID()] = set[test.From().ID()]
			} else {
				if(set[test.To().ID()] > set[test.From().ID()]) {
					set[test.From().ID()] = set[test.To().ID()]
				} else {
					set[test.To().ID()] = set[test.From().ID()]
				}
			}

			fmt.Println(nodeMap[int(test.From().ID())] + ":" + strconv.Itoa(int(set[test.From().ID()])) + " " + nodeMap[int(test.To().ID())] + ":" + strconv.Itoa(int(set[test.To().ID()])))
			fmt.Println("Adding Edge: " + nodeMap[int(test.From().ID())] + "->" + nodeMap[int(test.To().ID())])
			
			weight, _ := g.Weight(test.From(), test.To())
			mst.SetWeightedEdge(simple.WeightedEdge{
				F: simple.Node(test.From().ID()), 
				T: simple.Node(test.To().ID()), 
				W: weight})
		} else {
			fmt.Println("Not Adding Edge: " + strconv.Itoa(int(test.From().ID())) + "->" + strconv.Itoa(int(test.To().ID())))
		}
	}

	return mst
}

