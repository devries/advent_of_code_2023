package day21p2

import (
	"fmt"
	"io"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	return solveSteps(r, 26501365)
}

// 3456789
// 2#6789       +2
// 1#56789      +2
// 0#456789     +2
// 123456789
// 2#456789
// 3456789

func solveSteps(r io.Reader, steps int) uint64 {
	lines := utils.ReadLines(r)

	garden := Maze{
		positions: make(map[utils.Point]rune),
		bfs:       utils.NewBFS[utils.Point](),
		steps:     0, // start without a step limit to get available squares
		width:     len(lines[0]),
		height:    len(lines),
	}

	fmt.Printf("%#v\n", garden)

	for j, ln := range lines {
		for i, r := range ln {
			p := utils.Point{X: i, Y: j}
			garden.positions[p] = r
			if r == 'S' {
				garden.start = p
			}
		}
	}

	// for j := 0; j < garden.width; j += 2 {
	// 	for i := 0; i < garden.height; i += 2 {
	// 		fmt.Printf("%c", garden.positions[utils.Point{X: i, Y: j}])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	// for j := 1; j < garden.width; j += 2 {
	// 	for i := 1; i < garden.height; i += 2 {
	// 		fmt.Printf("%c", garden.positions[utils.Point{X: i, Y: j}])
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	// Get number of even and odd squares
	_, err := garden.bfs.Run(garden)
	if err != utils.BFSNotFound {
		panic("this should end when points are exhausted")
	}

	var even uint64
	var odd uint64

	for _, d := range garden.bfs.Distance {
		if d%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	var count uint64
	// How many maps fit within the total steps from S?
	fmt.Println("start:", garden.start, "width:", garden.width, "height", garden.height)
	fmt.Println("even:", even, "odd:", odd)

	mapNumber := (steps - garden.start.X) / garden.width

	additionalStepsDirect := steps - garden.start.X - mapNumber*garden.width
	additionalStepsDiag := steps - garden.start.X + garden.start.Y - (mapNumber-1)*garden.width
	fmt.Println("Maps:", mapNumber, "Extra:", additionalStepsDirect, "Extra diag:", additionalStepsDiag)

	evenmaps := 1
	oddmaps := 0
	for i := 1; i <= mapNumber; i++ {
		amt := i / 2
		if i%2 == 0 {
			evenmaps += 8 * amt
		} else {
			oddmaps += 8*amt + 4
		}
	}

	fmt.Println(evenmaps, oddmaps, even*uint64(evenmaps)+odd*uint64(oddmaps))
	// sp := garden.GetInitial()
	// for j := 0; j < garden.width; j++ {
	// 	for i := 0; i < garden.height; i++ {
	// 		pos := utils.Point{X: i, Y: j}
	// 		if pos == sp {
	// 			fmt.Printf("S")
	// 			continue
	// 		}
	// 		if garden.positions[pos] == '.' || garden.positions[pos] == 'S' {
	// 			if stp, ok := garden.bfs.Distance[pos]; ok {
	// 				stp = garden.bfs.Distance[pos]
	// 				expected := uint64(sp.Add(pos.Scale(-1)).Manhattan())
	// 				if expected != stp {
	// 					fmt.Printf("%d", stp-expected)
	// 				} else {
	// 					fmt.Printf(" ")
	// 				}
	// 			} else {
	// 				fmt.Printf("X")
	// 			}
	// 		} else {
	// 			fmt.Printf("#")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	// var count uint64
	// for _, d := range garden.bfs.Distance {
	// 	if d%2 == 0 {
	// 		count++
	// 	}
	// }

	return count
}

type Maze struct {
	positions map[utils.Point]rune
	bfs       *utils.BreadthFirstSearch[utils.Point]
	steps     int
	width     int
	height    int
	start     utils.Point
}

func (m Maze) GetInitial() utils.Point {
	return m.start
}

func (m Maze) GetNeighbors(pos utils.Point) []utils.Point {
	ret := []utils.Point{}

	// if m.bfs.Distance[pos] < uint64(m.steps) {
	// 	for _, dir := range utils.Directions {
	// 		np := pos.Add(dir)
	// 		npmapped := utils.Point{X: np.X % m.width, Y: np.Y % m.height}
	// 		if npmapped.X < 0 {
	// 			npmapped.X += m.width
	// 		}
	// 		if npmapped.Y < 0 {
	// 			npmapped.Y += m.height
	// 		}

	// 		if m.positions[npmapped] == '.' || m.positions[npmapped] == 'S' {
	// 			ret = append(ret, np)
	// 		}
	// 	}
	// }

	// Let's get distance to every point in map
	if m.steps == 0 || m.bfs.Distance[pos] < uint64(m.steps) {
		for _, dir := range utils.Directions {
			np := pos.Add(dir)
			if m.positions[np] == '.' || m.positions[np] == 'S' {
				ret = append(ret, np)
			}
		}
	}
	return ret
}

func (m Maze) IsFinal(pos utils.Point) bool {
	return false
}
