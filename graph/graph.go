// Implements an adjacency list graph as a slice of generic nodes
// and includes some useful graph functions.
package graph

import (
	"errors"
)

// Directed or undirected.
type GraphType int

const (
	Undirected GraphType = iota
	Directed
)

// Graph is an adjacency slice representation of a graph. Can be directed or undirected.
type Graph struct {
	nodes []*node
	Kind  GraphType
}

// New creates and returns an empty graph.
// If kind is Directed, returns a directed graph.
// This function returns an undirected graph by default.
func New(kind GraphType) *Graph {
	g := &Graph{}
	if kind == Directed {
		g.Kind = Directed
	}
	return g
}

// MakeNode creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	newNode := &node{Index: len(g.nodes)}
	newNode.container = Node{node: newNode, Value: new(interface{})}
	g.nodes = append(g.nodes, newNode)
	return newNode.container
}

// RemoveNode removes a node from the graph and all edges connected to it.
// This function nils points in the Node structure. If 'remove' is used in
// a map, you must delete the map index first.
func (g *Graph) RemoveNode(remove *Node) {
	if remove.node == nil {
		return
	}
	// O(V)
	nodeExists := false
	// remove all edges that connect from a different node to this one
	for _, node := range g.nodes {
		if node == remove.node {
			nodeExists = true
			continue
		}

		// O(E)
		swapIndex := -1 // index that the edge-to-remove is at
		for i := range node.edges {
			if node.edges[i].End == remove.node {
				swapIndex = i
			}
		}
		if swapIndex > -1 {
			swapNRemoveEdge(swapIndex, &node.edges)
		}

		// deal with possible reversed edge
		swapIndex = -1
		for i := range node.reversedEdges {
			if node.reversedEdges[i].End == remove.node {
				swapIndex = i
			}
		}
		if swapIndex > -1 {
			swapNRemoveEdge(swapIndex, &node.reversedEdges)
		}
		if node.Index > remove.node.Index {
			node.Index--
		}
	}
	if nodeExists {
		copy(g.nodes[remove.node.Index:], g.nodes[remove.node.Index+1:])
		g.nodes = g.nodes[:len(g.nodes)-1]
	}
	remove.node.parent = nil
	remove.node = nil
}

// MakeEdge calls MakeEdgeWeight with a weight of 0 and returns an error if either of the nodes do not
// belong in the graph. Calling MakeEdge multiple times on the same nodes will not create multiple edges.
func (g *Graph) MakeEdge(from, to Node) error {
	return g.MakeEdgeWeight(from, to, 0)
}

// MakeEdgeWeight creates  an edge in the graph with a corresponding weight.
// It returns an error if either of the nodes do not belong in the graph.
//
// Calling MakeEdgeWeight multiple times on the same nodes will not create multiple edges;
// this function will update the weight on the node to the new value.
func (g *Graph) MakeEdgeWeight(from, to Node, weight int) error {
	if from.node == nil || from.node.Index >= len(g.nodes) || g.nodes[from.node.Index] != from.node {
		return errors.New("First node in MakeEdge call does not belong to this graph")
	}
	if to.node == nil || to.node.Index >= len(g.nodes) || g.nodes[to.node.Index] != to.node {
		return errors.New("Second node in MakeEdge call does not belong to this graph")
	}

	for i := range from.node.edges { // check if edge already exists
		if from.node.edges[i].End == to.node {
			from.node.edges[i].Weight = weight

			// If the graph is undirected, fix the to node's weight as well
			if g.Kind == Undirected && to != from {
				for j := range to.node.edges {
					if to.node.edges[j].End == from.node {
						to.node.edges[j].Weight = weight
					}
				}
			}
			return nil
		}
	}
	newEdge := edge{Weight: weight, End: to.node}
	from.node.edges = append(from.node.edges, newEdge)
	reversedEdge := edge{Weight: weight, End: from.node} // weight for undirected graph only
	if g.Kind == Directed {                              // reversed edges are only used in directed graph algorithms
		to.node.reversedEdges = append(to.node.reversedEdges, reversedEdge)
	}
	if g.Kind == Undirected && to != from {
		to.node.edges = append(to.node.edges, reversedEdge)
	}
	return nil
}

// RemoveEdge removes edges starting at the from node and ending at the to node.
// If the graph is undirected, RemoveEdge will remove all edges between the nodes.
func (g *Graph) RemoveEdge(from, to Node) {
	fromEdges := from.node.edges
	toEdges := to.node.edges
	toReversedEdges := to.node.reversedEdges
	for e := range fromEdges { // fix from->to
		if fromEdges[e].End == to.node {
			swapNRemoveEdge(e, &fromEdges)
			from.node.edges = fromEdges
			break
		}
	}
	for e := range toReversedEdges { // fix reversed edges to->from
		if toReversedEdges[e].End == from.node {
			swapNRemoveEdge(e, &toReversedEdges)
			to.node.reversedEdges = toReversedEdges
			break
		}
	}
	if g.Kind == Undirected && from.node != to.node {
		for e := range toEdges {
			if toEdges[e].End == from.node {
				swapNRemoveEdge(e, &toEdges)
				to.node.edges = toEdges
				break
			}
		}
	}
}

// Neighbors returns a slice of nodes that are reachable from the given node in a graph.
func (g *Graph) Neighbors(n Node) []Node {
	neighbors := make([]Node, 0, len(n.node.edges))
	if g.nodes[n.node.Index] == n.node {
		for _, edge := range n.node.edges {
			neighbors = append(neighbors, edge.End.container)
		}
	}
	return neighbors
}

func (g *Graph) FetchNodes() []*node {
	return g.nodes
}

func (g *Graph) FetchEdges() []edge {
	var allEdges []edge
	for _, n := range g.nodes {
		for _, e := range n.edges {
			allEdges = append(allEdges, e)
		}
	}

	return allEdges
}

// Swaps an edge to the end of the edges slice and 'removes' it by reslicing.
func swapNRemoveEdge(remove int, edges *[]edge) {
	(*edges)[remove], (*edges)[len(*edges)-1] = (*edges)[len(*edges)-1], (*edges)[remove]
	*edges = (*edges)[:len(*edges)-1]
}