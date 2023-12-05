package day05p2

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
			// Build up the conversion array
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
			values = doConversion(conversions, values)

			conversions = []Conversion{}
		}
	}

	if len(conversions) > 0 {
		values = doConversion(conversions, values)
	}

	min := values[0]

	for i := 2; i < len(values); i += 2 {
		if values[i] < min {
			min = values[i]
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
// also get the range of subsequent numbers for which it is valid
func getDeltaInterval(arr []Conversion, val int64) (int64, int64) {
	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = low + (high-low)/2

		if arr[mid].Start <= val && arr[mid].End > val {
			return arr[mid].Delta, arr[mid].End - val
		}

		if arr[mid].Start > val {
			high = mid - 1
		} else if arr[mid].End <= val {
			low = mid + 1
		}
	}

	if arr[mid].Start > val {
		return 0, arr[mid].Start - val
	} else if mid < len(arr)-1 {
		return 0, arr[mid+1].Start - val
	} else {
		return 0, 0
	}
}

// Run conversions and return new values array
func doConversion(conversions []Conversion, values []int64) []int64 {
	// Conversion array complete, calculate conversions
	sort.Slice(conversions, func(i, j int) bool { return conversions[i].Start < conversions[j].Start })
	newvalues := []int64{}

	for i := 0; i < len(values); i += 2 {
		// For each input value and range we convert to a new value and range.
		// If the range is longer than the valid interval of the conversion we split up the range into two intervals
		// and then convert the second value and range as well... if that one is longer than the valid interval we repeat
		start, length := values[i], values[i+1]

		for {
			delta, interval := getDeltaInterval(conversions, start)
			newvalues = append(newvalues, start+delta)
			if length <= interval || interval == 0 { // 0 interval means the rest of the numbers follow that delta
				// The length of the input value range is less than the conversion interval
				newvalues = append(newvalues, length)
				break
			} else {
				// The length of the input value range is greater than the remaining
				// conversion interval, we need to split the solution up into multiple
				// ranges
				newvalues = append(newvalues, interval)
				start = start + interval
				length = length - interval
			}
		}
	}
	return newvalues
}
