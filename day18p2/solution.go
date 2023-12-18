package day18p2

import (
	"io"
	"regexp"
	"sort"
	"strconv"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	re := regexp.MustCompile(`([RDLU])\s+(\d+)\s+\(#([0-9a-f]+)\)`)

	pos := utils.Point{}
	xmin, ymin, xmax, ymax := 0, 0, 0, 0
	segments := SegmentList{}

	for _, ln := range lines {
		sm := re.FindStringSubmatch(ln)

		inst := Instruction{}
		actualDirection := []rune(sm[3])[5]
		switch actualDirection {
		case '3':
			inst.Direction = utils.North
		case '1':
			inst.Direction = utils.South
		case '0':
			inst.Direction = utils.East
		case '2':
			inst.Direction = utils.West
		default:
			panic("Direction not good")
		}

		distHex := string([]rune(sm[3])[:5])
		dist, err := strconv.ParseInt(distHex, 16, 32)
		if err != nil {
			utils.Check(err, "Unable to convert %s to int", sm[2])
		}
		inst.Distance = int(dist)

		newpos := pos.Add(inst.Direction.Scale(inst.Distance))

		// We'll be collecting horizontal segments and calculating
		// area by moving up from ymin to ymax.
		if inst.Direction == utils.East {
			segment := HSegment{pos, newpos}
			segments = append(segments, segment)
		} else if inst.Direction == utils.West {
			segment := HSegment{newpos, pos}
			segments = append(segments, segment)
		}
		pos = newpos

		if pos.X > xmax {
			xmax = pos.X
		}
		if pos.Y > ymax {
			ymax = pos.Y
		}
		if pos.X < xmin {
			xmin = pos.X
		}
		if pos.Y < ymin {
			ymin = pos.Y
		}
	}

	// There has got to be a simpler way to do whatever it is I do below, but I
	// was just not getting it today.

	// Sort all line segments by y and x, so we iterate from
	// lower left to upper right.
	sort.Sort(segments)
	ycurrent := ymin

	// Current ranges are segments over which we are calculating a filled in area.
	currentRanges := []int{}

	// Last motion ranges are the ranges over which we last caulculated filled in areas.
	lastMotionRanges := []int{}

	// This is the last border region calculated, which may be discarded if there has
	// been no vertical motion.
	lastBorderArea := uint64(0)

	var area uint64
	for _, s := range segments {
		d := s.Left.Y - ycurrent - 1 // hight between end segments. Ends are calculated separately.
		if d > 0 {
			// There has been a change in Y so we calculate the area of rectangles
			// with the current segment
			for i := 0; i < len(currentRanges); i += 2 {
				area += uint64(d) * uint64(currentRanges[i+1]-currentRanges[i]+1)
			}
			ycurrent = s.Left.Y
			// Save the last segments used to calculate area
			lastMotionRanges = make([]int, len(currentRanges))
			copy(lastMotionRanges, currentRanges)
		} else {
			// We didn't move vertically, so the last border calculation we had
			// was not a complete border, remove it and recalculate with next
			// segment.
			area -= lastBorderArea
		}

		previousRanges := make([]int, len(currentRanges))
		copy(previousRanges, currentRanges)

		// Add line segment to current ranges and find areas where areas will be
		// closed off or opened up by sorting the segments and calculating the
		// remaining alternating segments.
		// current: x-------------x          x---------------x
		// added:         x---------x
		// final:   x-----x       x-x        x---------------x
		currentRanges = append(currentRanges, s.Left.X, s.Right.X)
		sort.Ints(currentRanges)

		// Need to remove duplicate points
		for i := 0; i < len(currentRanges)-1; i++ {
			if currentRanges[i] == currentRanges[i+1] {
				// need to remove duplicates
				currentRanges = append(currentRanges[:i], currentRanges[i+2:]...)
				i -= 1
			}
		}

		// Since the last time there was movement find all overlap to get points in overlap
		// region between the areas just added and the areas added in the next iteration
		// lastMovement: x-------------x          x---------------x
		// current:            x---------x
		// overlap:      x---------------x        x---------------x
		boundaries := BoundaryList{}

		// Use method where we add 1 for start of regions and subtract one for end
		// any time sum > 0 there is at least one segment, and when sum == 0 then
		// there is no segment in either set of segments here.
		for i := 0; i < len(lastMotionRanges); i += 2 {
			boundaries = append(boundaries, Boundary{lastMotionRanges[i], 1})
			boundaries = append(boundaries, Boundary{lastMotionRanges[i+1], -1})
		}

		for i := 0; i < len(currentRanges); i += 2 {
			boundaries = append(boundaries, Boundary{currentRanges[i], 1})
			boundaries = append(boundaries, Boundary{currentRanges[i+1], -1})
		}

		sort.Sort(boundaries)

		current := 0
		combinedRanges := []int{}

		for _, b := range boundaries {
			pos := b.Position
			current += b.Incrementor
			if current == 1 && b.Incrementor == 1 {
				combinedRanges = append(combinedRanges, pos)
			}
			if current == 0 && b.Incrementor == -1 {
				combinedRanges = append(combinedRanges, pos)
			}
		}
		lastBorderArea = 0
		for i := 0; i < len(combinedRanges); i += 2 {
			v := uint64(combinedRanges[i+1] - combinedRanges[i] + 1)
			// keep the area here, because if we have another segment in this row to change
			// then we will have to subtract this change and do the problem again.
			lastBorderArea += v
			area += v
		}
	}

	return area
}

type Instruction struct {
	Direction utils.Point
	Distance  int
}

type HSegment struct {
	Left  utils.Point
	Right utils.Point
}

type SegmentList []HSegment

func (s SegmentList) Len() int      { return len(s) }
func (s SegmentList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s SegmentList) Less(i, j int) bool {
	li := s[i].Left
	lj := s[j].Left

	if li.Y != lj.Y {
		return li.Y < lj.Y
	}

	return li.X < lj.X
}

type Boundary struct {
	Position    int
	Incrementor int
}

type BoundaryList []Boundary

func (s BoundaryList) Len() int      { return len(s) }
func (s BoundaryList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s BoundaryList) Less(i, j int) bool {
	return s[i].Position < s[j].Position
}
