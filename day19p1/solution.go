package day19p1

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

	var sum int64
	for _, p := range partList {
		accepted := workflows.accept(p, "in")
		if utils.Verbose {
			fmt.Printf("%#v %t\n", p, accepted)
		}
		if accepted {
			sum += p.x + p.m + p.a + p.s
		}
	}
	return sum
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
