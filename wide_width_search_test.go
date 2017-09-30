package golang_algorithms

import (
	"testing"
)

var (
	graph *Graph
)

func init() {
	graph = NewGraph()
	graph.Save("you", []string{"alice", "bob", "claire"})
	graph.Save("bob", []string{"anuj", "peggy"})
	graph.Save("alice", []string{"peggy"})
	graph.Save("claire", []string{"thom", "jonny"})
	graph.Save("anuj", []string{})
	graph.Save("peggy", []string{})
	graph.Save("thom", []string{})
	graph.Save("jonny", []string{})
}

func TestWideWidthSearchTrue(t *testing.T) {
	i := graph.Search("jonny")
	if i != true {
		t.Errorf("expected result true; got %v", i)
	}
}

func TestWideWidthSearchFalse(t *testing.T) {
	i := graph.Search("serhii")
	if i != false {
		t.Errorf("expected result false; got %v", i)
	}
}