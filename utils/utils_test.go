package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestGcd(t *testing.T) {
	var tests = []struct {
		a      int64
		b      int64
		answer int64
	}{
		{1, 1, 1},
		{10, 5, 5},
		{12, 16, 4},
	}

	for _, test := range tests {
		result := Gcd(test.a, test.b)
		if result != test.answer {
			t.Errorf("For values %d and %d calculated %d, expected %d", test.a, test.b, result, test.answer)
		}
	}
}

func TestLcm(t *testing.T) {
	var tests = []struct {
		a      int64
		b      int64
		answer int64
	}{
		{1, 1, 1},
		{10, 5, 10},
		{12, 16, 48},
	}

	for _, test := range tests {
		result := Lcm(test.a, test.b)
		if result != test.answer {
			t.Errorf("For values %d and %d calculated %d, expected %d", test.a, test.b, result, test.answer)
		}
	}
}

func TestCountBits(t *testing.T) {
	var tests = []struct {
		n    uint64
		bits int64
	}{
		{0b0, 0},
		{0b10, 1},
		{0b1011010110, 6},
	}

	for _, test := range tests {
		result := CountBits(test.n)
		if result != test.bits {
			t.Errorf("For bitfield %b and calculated %d bits, expected %d bits", test.n, result, test.bits)
		}
	}
}

func TestPoint(t *testing.T) {
	p := Point{0, 0}

	p2 := p.Add(North)
	p3 := p.Add(East)

	p4 := North.Add(South)

	p5 := Point{12, -11}

	if p2 != North {
		t.Errorf("0,0 + North should be North")
	}

	if p3 != East {
		t.Errorf("0,0 + East should be East")
	}

	if p4 != p {
		t.Errorf("North + South should be 0,0")
	}

	if p5.Manhattan() != 23 {
		t.Errorf("The manhattan distance of {12, -11} should be 23, but I got %d", p5.Manhattan())
	}
}

func TestCheck(t *testing.T) {
	Check(nil, "test no error") // This should not panic

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The Check function did not panic")
		}
	}()

	err := fmt.Errorf("generic error")

	Check(err, "test error")
}

func TestReadLines(t *testing.T) {
	test := struct {
		Input  string
		Output []string
	}{"one\ntwo\nthree\n", []string{"one", "two", "three"}}

	r := strings.NewReader(test.Input)
	result := ReadLines(r)

	for i, v := range result {
		if v != test.Output[i] {
			t.Errorf("Expected %v, got %v", test.Output, result)
		}
	}
}

func TestOCR(t *testing.T) {
	input := `###..###..###...##..###...##...##..####.
#..#.#..#.#..#.#..#.#..#.#..#.#..#.#....
#..#.###..#..#.#..#.#..#.#..#.#....###..
###..#..#.###..####.###..####.#.##.#....
#.#..#..#.#....#..#.#.#..#..#.#..#.#....
#..#.###..#....#..#.#..#.#..#..###.#....`

	wanted := "RBPARAGF"

	obtained := OCRLetters(input)

	if strings.Compare(wanted, obtained) != 0 {
		t.Errorf("Expected %s, got %s", wanted, obtained)
	}
}

// This graph satisfies both the BFS interface and Dijkstra interface
type BFSGraph struct{}

func (b BFSGraph) GetInitial() Point {
	return Point{X: 0, Y: 0}
}

func (b BFSGraph) GetNeighbors(p Point) []Point {
	return []Point{{p.X + 1, p.Y}, {p.X - 1, p.Y}, {p.X, p.Y + 1}, {p.X, p.Y - 1}}
}

// Give every step a distance of 2 in the Dijkstra search
func (b BFSGraph) GetEdges(p Point) []Edge[Point] {
	pts := b.GetNeighbors(p)

	var edges []Edge[Point]
	for _, np := range pts {
		e := Edge[Point]{np, 2}
		edges = append(edges, e)
	}

	return edges
}

// Starting at (0,0) the final point is (10,0)
// Should be 10 steps for BFS and 20 distance for Dijkstra
func (b BFSGraph) IsFinal(p Point) bool {
	if p.X == 10 && p.Y == 0 {
		return true
	}

	return false
}

func TestBFS(t *testing.T) {
	bfs := NewBFS[Point]()

	var m BFSGraph

	v, err := bfs.Run(m)
	if err != nil {
		t.Errorf("Error in BFS search: %s", err)
	}

	if bfs.Distance[v] != 10 {
		t.Errorf("Expected 10 steps, got %d steps", bfs.Distance[v])
	}
}

func TestDijkstra(t *testing.T) {
	dij := NewDijkstra[Point]()

	var m BFSGraph

	v, err := dij.Run(m)
	if err != nil {
		t.Errorf("Error in Dijkstra search: %s", err)
	}

	if dij.Distance[v] != 20 {
		t.Errorf("Expected graph distance of 20, got %d", dij.Distance[v])
	}
}
