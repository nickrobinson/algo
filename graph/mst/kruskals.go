package mst

import (
	"fmt"
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"github.com/emirpasic/gods/maps/treemap"
	"strconv"
)

func Kruskals(g *simple.WeightedUndirectedGraph) {
	set := make(map[int64]int64)
	pq := treemap.NewWithIntComparator()

	for _, n := range g.Nodes() {
		set[n.ID()] = n.ID()
	}

	for _, e := range g.Edges() {
		weight, _ := g.Weight(e.From(), e.To())
		pq.Put(int(weight), e)
	}

	iter := pq.Iterator()
	for iter.Next() {
		//key := iter.Key()
		edge := iter.Value().(graph.Edge)
		if(set[edge.From().ID()] != set[edge.To().ID()]) {
			set[edge.To().ID()] = set[edge.From().ID()]
			fmt.Println("Adding Edge: " + strconv.Itoa(int(edge.From().ID())) + "->" + strconv.Itoa(int(edge.To().ID())))
		} else {
			fmt.Println("Not Adding Edge: " + strconv.Itoa(int(edge.From().ID())) + "->" + strconv.Itoa(int(edge.To().ID())))
		}
		//fmt.Println("Min Weight: " + strconv.Itoa(key.(int)))
	}
}

