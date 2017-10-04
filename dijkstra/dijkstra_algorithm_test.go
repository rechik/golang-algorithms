package dijkstra

import (
	"math"
	"testing"
)

var (
	graph map[string]map[string]int
	costs map[string]int
	parents map[string]string

	processed map[string]bool
)

func init() {
	graph = make(map[string]map[string]int)

	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1

	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5

	graph["fin"] = make(map[string]int)


	costs = make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt32


	parents = make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	processed = make(map[string]bool)
}

func TestDijkstraAlgorithm (t *testing.T) {
	node := FindLowestCostNode(costs)

	for node != "" {

		cost := costs[node]
		neighbors := graph[node]

		for key, neighborCost := range neighbors {
			newCost := cost + neighborCost

			if (costs[key] > newCost) {
				costs[key] = newCost
				parents[key] = node
			}
		}

		processed[node] = true

		node = FindLowestCostNode(costs)
	}

	if costs["fin"] != 6 {
		t.Errorf("expected result 6; got %v", costs["fin"])
	}
}

func FindLowestCostNode(costs map[string]int) string {
	lowestCost := math.MaxInt32
	var lowestCostNode string

	for node, cost := range costs{
		_, exists := processed[node];

		if (cost < lowestCost && !exists) {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}
