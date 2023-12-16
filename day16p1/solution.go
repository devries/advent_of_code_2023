package day16p1

import (
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	c := NewContraption(lines)

	bfs := utils.NewBFS[BeamState]()

	_, err := bfs.Run(c)
	if err != utils.BFSNotFound {
		utils.Check(err, "Error before searching full map")
	}

	return len(c.Energized)
}

type BeamState struct {
	Position  utils.Point
	Direction utils.Point
}

type Contraption struct {
	Tiles     map[utils.Point]rune
	Energized map[utils.Point]bool
}

func NewContraption(lines []string) *Contraption {
	v := Contraption{make(map[utils.Point]rune), make(map[utils.Point]bool)}

	for j, ln := range lines {
		for i, r := range ln {
			v.Tiles[utils.Point{X: i, Y: -j}] = r
		}
	}

	return &v
}

func (c *Contraption) GetInitial() BeamState {
	return BeamState{utils.Point{X: 0, Y: 0}, utils.East}
}

func (c *Contraption) GetNeighbors(s BeamState) []BeamState {
	tile := c.Tiles[s.Position]

	switch tile {
	case 0:
		// off map
		return []BeamState{}

	case '.':
		// continue on
		c.Energized[s.Position] = true
		newstate := BeamState{s.Position.Add(s.Direction), s.Direction}
		return []BeamState{newstate}

	case '\\':
		c.Energized[s.Position] = true
		var newDirection utils.Point
		switch s.Direction {
		case utils.North:
			newDirection = utils.West
		case utils.East:
			newDirection = utils.South
		case utils.South:
			newDirection = utils.East
		case utils.West:
			newDirection = utils.North
		}
		newstate := BeamState{s.Position.Add(newDirection), newDirection}
		return []BeamState{newstate}

	case '/':
		c.Energized[s.Position] = true
		var newDirection utils.Point
		switch s.Direction {
		case utils.North:
			newDirection = utils.East
		case utils.East:
			newDirection = utils.North
		case utils.South:
			newDirection = utils.West
		case utils.West:
			newDirection = utils.South
		}
		newstate := BeamState{s.Position.Add(newDirection), newDirection}
		return []BeamState{newstate}

	case '-':
		c.Energized[s.Position] = true
		switch s.Direction {
		case utils.North, utils.South:
			return []BeamState{{s.Position.Add(utils.East), utils.East}, {s.Position.Add(utils.West), utils.West}}
		default:
			return []BeamState{{s.Position.Add(s.Direction), s.Direction}}
		}

	case '|':
		c.Energized[s.Position] = true
		switch s.Direction {
		case utils.East, utils.West:
			return []BeamState{{s.Position.Add(utils.North), utils.North}, {s.Position.Add(utils.South), utils.South}}
		default:
			return []BeamState{{s.Position.Add(s.Direction), s.Direction}}
		}

	default:
		panic("Unexpected tile found")
	}
}

func (c *Contraption) IsFinal(s BeamState) bool {
	return false
}
