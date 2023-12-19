package day19p2

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)
	workflows := make(workflowMap)
	partList := []part{}

	startParts := false
	for _, ln := range lines {
		if ln == "" {
			startParts = true
			continue
		}

		if startParts {
			p := parsePart(ln)
			partList = append(partList, p)
		} else {
			name, instructions := parseWorkflow(ln)
			workflows[name] = instructions
		}
	}

	pr := partRange{1, 4000, 1, 4000, 1, 4000, 1, 4000}
	count := workflows.count(pr, "in")
	return count
}

type operation int

const (
	none operation = iota
	lessthan
	greaterthan
)

type part struct {
	x int64
	m int64
	a int64
	s int64
}

func (p part) getAttribute(a string) int64 {
	switch a {
	case "x":
		return p.x
	case "m":
		return p.m
	case "a":
		return p.a
	case "s":
		return p.s
	default:
		return 0
	}
}

type partRange struct {
	xmin int64
	xmax int64
	mmin int64
	mmax int64
	amin int64
	amax int64
	smin int64
	smax int64
}

func (pr partRange) count() int64 {
	var res int64
	if pr.xmax > pr.xmin {
		res = pr.xmax - pr.xmin + 1
	} else {
		return 0
	}
	if pr.mmax > pr.mmin {
		res *= pr.mmax - pr.mmin + 1
	} else {
		return 0
	}
	if pr.amax > pr.amin {
		res *= pr.amax - pr.amin + 1
	} else {
		return 0
	}
	if pr.smax > pr.smin {
		res *= pr.smax - pr.smin + 1
	} else {
		return 0
	}

	return res
}

type rule struct {
	op        operation
	attribute string
	argument  int64
	result    string
}

func (r rule) evaluate(p part) string {
	input := p.getAttribute(r.attribute)
	switch r.op {
	case none:
		return r.result
	case lessthan:
		if input < r.argument {
			return r.result
		}
	case greaterthan:
		if input > r.argument {
			return r.result
		}
	default:
		panic("unknown operation")
	}

	return ""
}

// use rule to split part range into a true and false portion
func (r rule) splitRange(pr partRange) (partRange, partRange) {
	trueRange := pr
	falseRange := pr

	switch r.attribute {
	case "x":
		min, max := pr.xmin, pr.xmax
		trueRange.xmin, trueRange.xmax = rangeReduce(r.op, r.argument, min, max)
		falseRange.xmin, falseRange.xmax = rangeReduceOther(r.op, r.argument, min, max)
	case "m":
		min, max := pr.mmin, pr.mmax
		trueRange.mmin, trueRange.mmax = rangeReduce(r.op, r.argument, min, max)
		falseRange.mmin, falseRange.mmax = rangeReduceOther(r.op, r.argument, min, max)
	case "a":
		min, max := pr.amin, pr.amax
		trueRange.amin, trueRange.amax = rangeReduce(r.op, r.argument, min, max)
		falseRange.amin, falseRange.amax = rangeReduceOther(r.op, r.argument, min, max)
	case "s":
		min, max := pr.smin, pr.smax
		trueRange.smin, trueRange.smax = rangeReduce(r.op, r.argument, min, max)
		falseRange.smin, falseRange.smax = rangeReduceOther(r.op, r.argument, min, max)
	}

	return trueRange, falseRange
}

func rangeReduce(op operation, arg int64, min int64, max int64) (int64, int64) {
	switch op {
	case none:
		return min, max
	case lessthan:
		if max >= arg {
			max = arg - 1
		}
	case greaterthan:
		if min <= arg {
			min = arg + 1
		}
	default:
		panic("unknown operation")
	}

	return min, max
}

func rangeReduceOther(op operation, arg int64, min int64, max int64) (int64, int64) {
	switch op {
	case none:
		return 0, 0
	case lessthan:
		if min < arg {
			min = arg
		}
	case greaterthan:
		if max > arg {
			max = arg
		}
	default:
		panic("unknown operation")
	}

	return min, max
}

type workflowMap map[string][]rule

// Check if part is accepted starting with workflow start
func (w workflowMap) accept(p part, start string) bool {
	instructions := w[start]

	for _, i := range instructions {
		result := i.evaluate(p)
		switch result {
		case "A":
			return true
		case "R":
			return false
		case "":
			// do nothing and move on
		default:
			return w.accept(p, result)
		}
	}

	panic("Finished workflow with no result")
}

func (w workflowMap) count(pr partRange, start string) int64 {
	instructions := w[start]
	var trueRange, falseRange partRange

	var sum int64

	for _, i := range instructions {
		trueRange, falseRange = i.splitRange(pr)
		switch i.result {
		case "A":
			// Accepted range
			if utils.Verbose {
				fmt.Printf("Accepted: %#v\n", trueRange)
			}
			sum += trueRange.count()
		case "R":
			// no further action on trueRange
		case "":
			// not sure about this one
		default:
			sum += w.count(trueRange, i.result)
		}
		pr = falseRange
	}

	return sum
}

func parseWorkflow(ln string) (string, []rule) {
	rules := []rule{}

	ln = strings.TrimSuffix(ln, "}")
	parts := strings.Split(ln, "{")

	name := parts[0]

	ruleStatements := strings.Split(parts[1], ",")

	for _, s := range ruleStatements {
		ruleComponents := strings.Split(s, ":")

		if len(ruleComponents) == 1 {
			// This is a non rule
			r := rule{none, "", 0, ruleComponents[0]}
			rules = append(rules, r)
		} else {
			opRune := ruleComponents[0][1]
			var op operation
			switch opRune {
			case '>':
				op = greaterthan
			case '<':
				op = lessthan
			default:
				panic("unknown operation in parse")
			}

			arg, err := strconv.ParseInt(ruleComponents[0][2:], 10, 64)
			utils.Check(err, "Unable to parse %s to integer", ruleComponents[0][2:])
			r := rule{op, ruleComponents[0][:1], arg, ruleComponents[1]}
			rules = append(rules, r)
		}
	}

	return name, rules
}

func parsePart(ln string) part {
	ln = strings.TrimSuffix(ln, "}")
	ln = strings.TrimPrefix(ln, "{")

	parts := strings.Split(ln, ",")

	ret := part{}

	for _, p := range parts {
		sides := strings.Split(p, "=")
		value, err := strconv.ParseInt(sides[1], 10, 64)
		utils.Check(err, "unable to parse %s to integer", sides[1])

		switch sides[0] {
		case "x":
			ret.x = value
		case "m":
			ret.m = value
		case "a":
			ret.a = value
		case "s":
			ret.s = value
		}
	}

	return ret
}
