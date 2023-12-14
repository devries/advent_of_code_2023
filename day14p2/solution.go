package day14p2

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
	"io"
	"sort"

	"aoc/utils"
)

var maxiter = 1000000000

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	ymax := len(lines)
	xmax := len(lines[0])

	pos := make(map[utils.Point]rune)

	for j, ln := range lines {
		for i, r := range ln {
			if r != '.' {
				p := utils.Point{X: i, Y: ymax - j - 1}
				pos[p] = r
			}
		}
	}

	platform := Platform{xmax, ymax, pos}

	checksums := make(map[uint64]int)
	looking := true
	for i := 0; i < maxiter; i++ {
		tiltCycle(&platform)
		cs := platform.checkSum()
		if v, ok := checksums[cs]; ok && looking {
			interval := i - v
			start := v
			span := maxiter - start
			ncycles := span / interval
			i = start + ncycles*interval
			looking = false
		}
		checksums[cs] = i
	}

	if utils.Verbose {
		fmt.Println(platform.checkSum())
		platform.print()
	}
	// measure load
	sum := 0
	for k, v := range platform.Positions {
		if v == 'O' {
			sum += k.Y + 1
		}
	}
	return sum
}

type Platform struct {
	XMax      int
	YMax      int
	Positions map[utils.Point]rune
}

func tiltCycle(platform *Platform) {
	// Tilt north
	for i := 0; i < platform.XMax; i++ {
		blocker := platform.YMax // what blocks the stones
		for j := platform.YMax - 1; j >= 0; j-- {
			pt := utils.Point{X: i, Y: j}
			r := platform.Positions[pt]
			switch r {
			case 'O':
				newPt := utils.Point{X: i, Y: blocker - 1}
				delete(platform.Positions, pt)
				platform.Positions[newPt] = 'O'
				blocker = blocker - 1
			case '#':
				blocker = j
			}
		}
	}

	// Tilt west
	for j := 0; j < platform.YMax; j++ {
		blocker := -1 // what blocks the stones
		for i := 0; i < platform.XMax; i++ {
			pt := utils.Point{X: i, Y: j}
			r := platform.Positions[pt]
			switch r {
			case 'O':
				newPt := utils.Point{X: blocker + 1, Y: j}
				delete(platform.Positions, pt)
				platform.Positions[newPt] = 'O'
				blocker = blocker + 1
			case '#':
				blocker = i
			}
		}
	}

	// Tilt south
	for i := 0; i < platform.XMax; i++ {
		blocker := -1 // what blocks the stones
		for j := 0; j < platform.YMax; j++ {
			pt := utils.Point{X: i, Y: j}
			r := platform.Positions[pt]
			switch r {
			case 'O':
				newPt := utils.Point{X: i, Y: blocker + 1}
				delete(platform.Positions, pt)
				platform.Positions[newPt] = 'O'
				blocker = blocker + 1
			case '#':
				blocker = j
			}
		}
	}

	// Tilt east
	for j := 0; j < platform.YMax; j++ {
		blocker := platform.XMax // what blocks the stones
		for i := platform.XMax - 1; i >= 0; i-- {
			pt := utils.Point{X: i, Y: j}
			r := platform.Positions[pt]
			switch r {
			case 'O':
				newPt := utils.Point{X: blocker - 1, Y: j}
				delete(platform.Positions, pt)
				platform.Positions[newPt] = 'O'
				blocker = blocker - 1
			case '#':
				blocker = i
			}
		}
	}
}

func (p *Platform) print() {
	for j := p.YMax - 1; j >= 0; j-- {
		for i := 0; i < p.XMax; i++ {
			pt := utils.Point{X: i, Y: j}
			r := p.Positions[pt]
			if r == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (p *Platform) checkSum() uint64 {
	parray := []int{}
	for j := p.YMax - 1; j >= 0; j-- {
		for i := 0; i < p.XMax; i++ {
			pt := utils.Point{X: i, Y: j}
			r := p.Positions[pt]
			if r == 'O' {
				parray = append(parray, p.XMax*j+i)
			}
		}
	}

	sort.Ints(parray)
	hf := fnv.New64a()

	for _, v := range parray {
		err := binary.Write(hf, binary.LittleEndian, int32(v))
		utils.Check(err, "Error writing binary")
	}

	return hf.Sum64()
}
