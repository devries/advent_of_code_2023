package day20p2

import (
	"fmt"
	"io"
	"os"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	modules := make(map[string]*module)
	for _, ln := range lines {
		name, mod := parseLine(ln)
		modules[name] = mod
	}

	// get all sources for modules
	for name, mod := range modules {
		for _, d := range mod.destinations {
			if md, ok := modules[d]; ok {
				md.sources = append(md.sources, name)
			}
		}
	}

	f, err := os.Create("day20p2/relations.gv")
	utils.Check(err, "unable to open visualization file")
	defer f.Close()

	// write out a graphviz visualization
	fmt.Fprintln(f, "digraph G {")

	for name, m := range modules {
		var shape string
		switch m.kind {
		case flipflop:
			shape = "rect"
		case conjunction:
			shape = "trapezium"
		default:
			shape = "circle"
		}

		fmt.Fprintf(f, "  %s [ shape = \"%s\"; ];\n", name, shape)
		for _, d := range m.destinations {
			fmt.Fprintf(f, "  %s -> %s;\n", name, d)
		}
	}
	fmt.Fprintln(f, "}")

	// Can deduce that there are 12 bit counters that trigger an and LOW
	// and resets the clock. When all ands are low at the same time you get
	// your pulse. The counter trip point is when all flip-flops that feed
	// into conjunction gate are HIGH, and the ones which receive a signal
	// from that AND gate are LOW.
	//
	// The lSB is next to the button, and the bit order follows the links
	// between flip-flops.

	val := int64(0b111101011011)                // kb
	val = utils.Lcm(val, int64(0b111100010111)) // vm
	val = utils.Lcm(val, int64(0b111011010101)) // dn
	val = utils.Lcm(val, int64(0b111010111001)) // vk

	return val
}

type modtype int

const (
	flipflop modtype = iota
	conjunction
	broadcast
)

type module struct {
	kind         modtype
	destinations []string
	sources      []string
	bitnumber    int
}

func parseLine(ln string) (string, *module) {
	var ret module
	var name string

	components := strings.Split(ln, " -> ")
	ret.destinations = strings.Split(components[1], ", ")

	switch components[0][0] {
	case '%':
		// flip flop
		ret.kind = flipflop
		name = components[0][1:]
	case '&':
		ret.kind = conjunction
		name = components[0][1:]
	default:
		ret.kind = broadcast
		name = components[0]
	}

	return name, &ret
}
