package trip

import (
"fmt"
"github.com/starwander/goraph"
)

type myVertex struct {
	name     string
	time 	 string
	outTo  map[string]float64
	inFrom map[string]float64
}

func (vertex *myVertex) ID() goraph.ID {
	panic("implement me")
}

type myEdge struct {
	from   string
	to     string
	weight float64
}

func (vertex *myVertex) Id() goraph.ID {
	return vertex.name
}

func (vertex *myVertex) Edges() (edges []goraph.Edge) {
	edges = make([]goraph.Edge, len(vertex.outTo)+len(vertex.inFrom))
	i := 0
	for to, weight := range vertex.outTo {
		edges[i] = &myEdge{vertex.name, to, weight}
		i++
	}
	for from, weight := range vertex.inFrom {
		edges[i] = &myEdge{from, vertex.name, weight}
		i++
	}
	return
}

func (edge *myEdge) Get() (goraph.ID, goraph.ID, float64) {
	return edge.from, edge.to, edge.weight
}

func main() {
	//Dikjstra: The distance from S to T is  44
	//Dikjstra: The path from S to T is: T<-F<-E<-B<-S
	Dijkstra()

	//Yen 1st: The distance from A to D is  2
	//Yen 1st: The path from A to D is:  [A D]
	//Yen 2nd: The distance from A to D is  3
	//Yen 2nd: The path from A to D is:  [A B C D]
	//Yen 3rd: The distance from A to D is  4
	//Yen 3rd: The path from A to D is:  [A B E D]
	Yen()

	//Kisp 1st: The distance from A to D is  5
	//Kisp 1st: The path from A to D is:  [C E F H]
	//Kisp 2nd: The distance from A to D is  11
	//Kisp 2nd: The path from A to D is:  [C D F G H]
	//Kisp 3rd: The distance from A to D is  +Inf
	//Kisp 3rd: The path from A to D is:  []
	Kisp()
}

func Dijkstra() {
	graph := goraph.NewGraph()
	//_ = graph.AddVeges(&myVertex{"T", map[string]float64{"A": 44, "D": 16, "E": 19, "F": 6}, map[string]float64{"A": 44, "D": 16, "E": 19, "F": 6}})

	dist, prev, _ := graph.Dijkstra("S")
	fmt.Println("Dikjstra: The distance from S to T is ", dist["T"])
	fmt.Print("Dikjstra: The path from S to T is: T")
	node := prev["T"]
	for node != nil {
		fmt.Printf("<-%s", node)
		node = prev[node]
	}
	fmt.Println()
}

func Yen() {
	graph := goraph.NewGraph()
	//graph.AddVertex("A", nil)
	//graph.AddVertex("B", nil)
	//graph.AddVertex("C", nil)
	//graph.AddVertex("D", nil)
	//graph.AddVertex("E", nil)
	//graph.AddEdge("A", "B", 1, nil)
	//graph.AddEdge("B", "C", 1, nil)
	//graph.AddEdge("C", "D", 1, nil)
	//graph.AddEdge("A", "D", 2, nil)
	//graph.AddEdge("B", "E", 2, nil)
	//graph.AddEdge("E", "D", 1, nil)
	dist, path, _ := graph.Yen("A", "D", 3)
	fmt.Println("Yen 1st: The distance from A to D is ", dist[0])
	fmt.Println("Yen 1st: The path from A to D is: ", path[0])
	fmt.Println("Yen 2nd: The distance from A to D is ", dist[1])
	fmt.Println("Yen 2nd: The path from A to D is: ", path[1])
	fmt.Println("Yen 3rd: The distance from A to D is ", dist[2])
	fmt.Println("Yen 3rd: The path from A to D is: ", path[2])
}

func Kisp() {
	graph := goraph.NewGraph()
	//graph.AddVertexWithEdges(&myVertex{"C", map[string]float64{"D": 3, "E": 2}, map[string]float64{}})
	//graph.AddVertexWithEdges(&myVertex{"D", map[string]float64{"F": 4}, map[string]float64{"C": 3, "E": 1}})
	//graph.AddVertexWithEdges(&myVertex{"E", map[string]float64{"D": 1, "F": 2, "G": 3}, map[string]float64{"C": 2}})
	//graph.AddVertexWithEdges(&myVertex{"F", map[string]float64{"G": 2, "H": 1}, map[string]float64{"D": 4, "E": 2}})
	//graph.AddVertexWithEdges(&myVertex{"G", map[string]float64{"H": 2}, map[string]float64{"E": 3, "F": 2}})
	//graph.AddVertexWithEdges(&myVertex{"H", map[string]float64{}, map[string]float64{"F": 1, "G": 2}})
	dist, path, _ := graph.Kisp("C", "H", 3)
	fmt.Println("Kisp 1st: The distance from A to D is ", dist[0])
	fmt.Println("Kisp 1st: The path from A to D is: ", path[0])
	fmt.Println("Kisp 2nd: The distance from A to D is ", dist[1])
	fmt.Println("Kisp 2nd: The path from A to D is: ", path[1])
	fmt.Println("Kisp 3rd: The distance from A to D is ", dist[2])
	fmt.Println("Kisp 3rd: The path from A to D is: ", path[2])
}

func fillGraph(waypoints []string, waypointsTimes []string)  {
	//graph := goraph.NewGraph()

	//for i := 0; i < len(waypoints) ; i++  {
	//	otherNodes := make(map[string]int)
	//	for j := range waypoints {
	//		if j != i {
	//			otherNodes[waypoints[j]] = 10
	//			otherNodes[waypoints[j]] = 10
	//		}
	//	}
	//	_ = graph.AddVertexWithEdges(&myVertex{waypoints[i], waypointsTimes[i], map[string]float64{"B": 14}, map[string]float64{"A": 15, "B": 14, "C": 9}})
	//	graph.
	//	for j := 0; j < len(waypoints) ; j++ {
	//		if j != i {
	//			g.AddEdge(&node, &Node{waypoints[j], waypointsTime[j]} )
	//		}
	//
	//	}
	//}
	//
	//return graph
}
