package golang_algorithms

import (
	"container/list"
)

// Describes structure of graph
type Graph struct {
	values map[string][]string
	queue *list.List
	searched map[string]bool
}

// Put value to graph
func (g *Graph) Save(key string, value []string)  {
	g.values[key] = value
}

// Get value from graph
func (g *Graph) Load(key string) []string {
	return g.values[key]
}

// Search nearest friend
func (g *Graph) Search(name string) bool {
	g.pushFriendsToQueue(g.Load("you"))

	for e := g.queue.Front(); e != nil; e = e.Next() {
		if _, exist := g.searched[name]; !exist {
			if name == e.Value {
				return true
			} else {
				g.pushFriendsToQueue(g.Load(e.Value.(string)))
			}
		}
	}

	return false
}

// Internal method. Put friends to checking
func (g *Graph) pushFriendsToQueue(friends []string)  {
	for _, v := range friends {
		g.queue.PushBack(v)
	}
}

// Constructor
func NewGraph() *Graph {
	return &Graph{
		values: make(map[string][]string),
		queue: list.New(),
		searched: make(map[string]bool),
	}
}