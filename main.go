package main

//go:generate go run ./gen

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"aoc/utils"
	"github.com/spf13/pflag"
)

var RunAll bool
var DaySelected int
var PartSelected string

func init() {
	pflag.BoolVarP(&RunAll, "all", "a", false, "run all days")
	pflag.IntVarP(&DaySelected, "day", "d", 0, "run specific day")
	pflag.StringVarP(&PartSelected, "part", "p", "2", "run specific part (default 2)")
}

func main() {
	pflag.Parse()
	if RunAll {
		runAll()
		return
	}

	switch DaySelected {
	case 0:
		runCurrent()
	default:
		runDay(DaySelected, PartSelected)
	}
}

type aocFunc func(io.Reader) any

type aocResult struct {
	Result      string
	TimeElapsed time.Duration
}

type aocRunnerInput struct {
	Name     string
	Func     aocFunc
	Filename string
	Day      int
	Part     string
}

func runAocPart(partFunc aocFunc, filename string) aocResult {
	f, err := os.Open(filename)
	utils.Check(err, "unable to open file %s", filename)
	defer f.Close()

	start := time.Now()
	r := partFunc(f)
	duration := time.Since(start)

	res := aocResult{TimeElapsed: duration}

	switch v := r.(type) {
	case int:
		res.Result = strconv.Itoa(v)
	case int64:
		res.Result = strconv.FormatInt(v, 10)
	case uint64:
		res.Result = strconv.FormatUint(v, 10)
	case string:
		res.Result = v
	case fmt.Stringer:
		res.Result = v.String()
	default:
		res.Result = "unknown return value"
	}

	return res
}

func runAll() {
	var r aocResult
	var total time.Duration

	for _, v := range days {
		r = runAocPart(v.Func, v.Filename)
		total += r.TimeElapsed

		fmt.Printf("%s: %s time elapsed: %s\n", v.Name, r.Result, r.TimeElapsed)
	}

	fmt.Printf("Overall time elapsed: %s\n", total)
}

func runDay(day int, part string) {
	found := false

	for _, v := range days {
		if v.Day == day && v.Part == part {
			fmt.Printf("Day %d part %s\n", day, part)
			r := runAocPart(v.Func, v.Filename)
			fmt.Println(r.Result)
			fmt.Printf("Time elapsed: %s\n", r.TimeElapsed)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Did not find a solution for day %d part %s\n", day, part)
	}
}
