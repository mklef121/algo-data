
package main

import (
    "errors"
    "fmt"
)

type Node struct {
	x int
	y int
	weight float64
	heuristic float64
	distance float64
	neighbors []Node
}

type Graph struct {
	nodes []Node
}

func InitNode(_x int, _y int, _weight float64) Node {
    n := Node {
        x: _x,
        y: _y,
        weight: _weight,
		heuristic: 0.0,
		distance: 0.0,
    }
    return n
}


func InitGraph() Graph {
	graph := Graph{}
	return graph
}

func AddNodeToGraph(graph Graph, node Node) {
    graph.nodes = append(graph.nodes, node)
}

func FindNodeInGraph(graph Graph, x int, y int) (*Node, error) {
	for i := 0; i < len(graph.nodes); i++ {
		if graph.nodes[i].x == x && graph.nodes[i].y == y {
			return &graph.nodes[i], nil
        }
	}

    return nil, errors.New("empty name")
}

func AddEdge(graph Graph, xNode1 int, yNode1 int, xNode2 int, yNode2 int) (bool, error) {
    var node1 *Node
    var err error
    node1, err = FindNodeInGraph(graph, xNode1, yNode1);
    fmt.Println(node1)
    return true, err
}

func main() {
    var graph Graph
    var nodes []Node
    var maxX = 10
    var maxY = 10
    for x := 0; x < maxX; x++ {
        for y := 0; y < maxY; y++ {
            node := InitNode(x, y, 1.0)
            nodes = append(nodes, node)
            AddNodeToGraph(graph, node)
        }
    }
    

    fmt.Printf("Created %d nodes in tree\n", maxX * maxY)
}
