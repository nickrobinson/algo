package graph

type edge struct {
	Weight int
	End    *node
}

// An Edge connects two Nodes in a graph. To modify Weight, use
// the MakeEdgeWeight function. Any local modifications will
// not be seen in the graph.
type Edge struct {
	Weight int
	Start  Node
	End    Node
}