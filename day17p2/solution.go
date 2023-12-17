package day17p2

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	m := NewMaze(lines)

	dij := utils.NewDijkstra[State]()

	finalState, err := dij.Run(m)
	utils.Check(err, "error in search")

	return dij.Distance[finalState]
}

type Maze struct {
	HeatLoss    map[utils.Point]uint64
	BottomRight utils.Point
}

func NewMaze(lines []string) *Maze {
	m := Maze{HeatLoss: make(map[utils.Point]uint64)}

	for j, ln := range lines {
		for i, r := range ln {
			p := utils.Point{X: i, Y: -j}
			v := uint64(r - '0')
			m.HeatLoss[p] = v
			m.BottomRight = p
		}
	}

	return &m
}

type State struct {
	Position  utils.Point
	Direction utils.Point
	Steps     int
}

func (m *Maze) GetInitial() State {
	return State{utils.Point{X: 0, Y: 0}, utils.East, 0}
}

func (m *Maze) GetEdges(s State) []utils.Edge[State] {
	ret := []utils.Edge[State]{}

	// forward
	np := s.Position.Add(s.Direction)
	if hl, ok := m.HeatLoss[s.Position.Add(s.Direction)]; ok && s.Steps < 10 {
		newState := State{np, s.Direction, s.Steps + 1}
		ret = append(ret, utils.Edge[State]{Node: newState, Distance: hl})
	}

	// Left and Right
	for _, dir := range []utils.Point{s.Direction.Right(), s.Direction.Left()} {
		np := s.Position.Add(dir)
		if hl, ok := m.HeatLoss[np]; ok && s.Steps >= 4 {
			newState := State{np, dir, 1}
			ret = append(ret, utils.Edge[State]{Node: newState, Distance: hl})
		}
	}

	return ret
}

func (m *Maze) IsFinal(s State) bool {
	if s.Position == m.BottomRight && s.Steps >= 4 {
		return true
	}

	return false
}
