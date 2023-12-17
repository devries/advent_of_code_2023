package utils

import (
	"container/heap"
	"errors"
)

var DijkstraNotFound = errors.New("dijkstra: unable to find final state")

type DijkstraSearch[T comparable] struct {
	Queue    []T
	Visited  map[T]bool
	Previous map[T]T
	Distance map[T]uint64
	Index    map[T]int
}

// This interface encompases the functionality of a graph. It can
// provide the initial state, edges for each state, and identify
// if a state is a final state of the system.
type WeightedGraph[T comparable] interface {
	// Return the initial node in the graph
	GetInitial() T

	// Get the edges attached from a particular node
	GetEdges(T) []Edge[T]

	// True of the node is a valid final state of the problem
	IsFinal(T) bool
}

// Edge is a structure which holds both a node and the distance to that node.
// A WeightedGraph will return an array of edges from a given node to the neighboring
// nodes.
type Edge[T comparable] struct {
	Node     T
	Distance uint64
}

// Prepare a new Dijkstra search structure for notes of type T.
// T must be comparable and will be used as a map key.
func NewDijkstra[T comparable]() *DijkstraSearch[T] {
	s := DijkstraSearch[T]{[]T{}, make(map[T]bool), make(map[T]T), make(map[T]uint64), make(map[T]int)}
	return &s
}

// ---- HEAP DEFINITIONS ----
func (b *DijkstraSearch[T]) Len() int { return len(b.Queue) }

func (b *DijkstraSearch[T]) Less(i, j int) bool {
	return b.Distance[b.Queue[i]] < b.Distance[b.Queue[j]]
}

func (b *DijkstraSearch[T]) Swap(i, j int) {
	b.Queue[i], b.Queue[j] = b.Queue[j], b.Queue[i]
	b.Index[b.Queue[i]] = i
	b.Index[b.Queue[j]] = j
}

func (b *DijkstraSearch[T]) Pop() any {
	old := b.Queue
	n := len(old)
	item := old[n-1]
	b.Index[item] = -1
	b.Queue = old[0 : n-1]
	return item
}

func (b *DijkstraSearch[T]) Push(x any) {
	item := x.(T)
	n := len(b.Queue)
	b.Index[item] = n
	b.Queue = append(b.Queue, item)
}

// Push or update a new node onto the Dijkstra queue along with the
// distance and previous node
func (b *DijkstraSearch[T]) PushOrUpdate(s T, d uint64, previous T) {
	idx, ok := b.Index[s]
	if ok {
		// seen before
		if d >= b.Distance[s] {
			// not a better distance, so do nothing
			return
		}
		b.Distance[s] = d
		b.Previous[s] = previous
		heap.Fix(b, idx)
	} else {
		// never seen before, so push it
		b.Distance[s] = d
		b.Previous[s] = previous
		heap.Push(b, s)
	}
}

// --------------------------

// Run a dijkstra search returning the final state or an error if the final state was not
// reachable.
func (b *DijkstraSearch[T]) Run(g WeightedGraph[T]) (T, error) {
	initState := g.GetInitial()
	b.Distance[initState] = 0
	b.Queue = append(b.Queue, initState)
	heap.Init(b) // do heap

	// Do bfs
	for len(b.Queue) > 0 {
		s := heap.Pop(b).(T)
		d := b.Distance[s]

		if g.IsFinal(s) {
			return s, nil
		}

		if b.Visited[s] {
			continue
		}
		b.Visited[s] = true

		for _, edge := range g.GetEdges(s) {
			ns := edge.Node
			delta := edge.Distance

			if b.Visited[ns] {
				continue
			}

			totald := d + delta
			b.PushOrUpdate(ns, totald, s)
		}
	}

	return initState, DijkstraNotFound
}

// GetPath returns the shortest from the starting node to the given node.
func (b *DijkstraSearch[T]) GetPath(s T) []T {
	ret := []T{s}

	for {
		ns, ok := b.Previous[s]
		if !ok {
			break
		}
		ret = append(ret, ns)
		s = ns
	}

	// reverse
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
