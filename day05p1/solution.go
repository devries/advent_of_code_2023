package day05p1

import (
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	// Get seeds from first line
	parts := strings.Fields(lines[0])

	values := make([]int64, len(parts)-1)

	for i := 0; i < len(values); i++ {
		v, err := strconv.ParseInt(parts[i+1], 10, 64)
		utils.Check(err, "Unable to convert %s to int64", parts[i+1])

		values[i] = v
	}

	// For remaining lines create conversions and convert values
	conversions := []Conversion{}

	// three positive integers
	re := regexp.MustCompile(`\d+\s+\d+\s+\d+`)

	for _, ln := range lines[2:] {
		switch {
		case re.MatchString(ln):
			parts = strings.Fields(ln)

			components := make([]int64, len(parts))

			for i, s := range parts {
				var err error
				components[i], err = strconv.ParseInt(s, 10, 64)
				utils.Check(err, "Unable to convert %s to int64", s)
			}
			c := Conversion{Start: components[1], End: components[1] + components[2], Delta: components[0] - components[1]}
			conversions = append(conversions, c)

		case ln == "":
			sort.Slice(conversions, func(i, j int) bool { return conversions[i].Start < conversions[j].Start })

			for i, v := range values {
				delta := getDelta(conversions, v)
				values[i] = v + delta
			}

			conversions = []Conversion{}
		}
	}

	if len(conversions) > 0 {
		sort.Slice(conversions, func(i, j int) bool { return conversions[i].Start < conversions[j].Start })

		for i, v := range values {
			delta := getDelta(conversions, v)
			values[i] = v + delta
		}
	}

	min := values[0]

	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

type Conversion struct {
	Start int64
	End   int64
	Delta int64
}

// Get the delta value by performing a binary search over sorted array of conversions
func getDelta(arr []Conversion, val int64) int64 {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		if arr[mid].Start <= val && arr[mid].End > val {
			return arr[mid].Delta
		}

		if arr[mid].Start > val {
			high = mid - 1
		} else if arr[mid].End <= val {
			low = mid + 1
		}
	}

	return 0
}
