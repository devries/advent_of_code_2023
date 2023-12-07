package day07p1

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"aoc/utils"
)

func Solve(r io.Reader) any {
	lines := utils.ReadLines(r)

	var g Game
	for _, ln := range lines {
		hand := NewHand(ln)
		g = append(g, hand)
	}

	sort.Sort(g)

	var sum int64
	for i, v := range g {
		sum += int64(i+1) * v.Bid
	}
	return sum
}

type HandType int64

const (
	HIGHCARD HandType = iota
	ONEPAIR
	TWOPAIR
	THREEOFAKIND
	FULLHOUSE
	FOUROFAKIND
	FIVEOFAKIND
)

type Hand struct {
	Type  HandType
	Cards []rune
	Bid   int64
}

// Create hand by parsing line
func NewHand(ln string) *Hand {
	parts := strings.Fields(ln)

	cards := []rune(parts[0])
	bid, err := strconv.ParseInt(parts[1], 10, 64)
	utils.Check(err, "Unable to parse %s to int64", parts[1])

	counts := make(map[rune]int)

	for _, c := range cards {
		counts[c]++
	}

	cardcounts := make([]int, 0, 5)

	for _, v := range counts {
		cardcounts = append(cardcounts, v)
	}

	sort.Ints(cardcounts)

	var hand HandType
	switch {
	case equal(cardcounts, []int{5}):
		hand = FIVEOFAKIND
	case equal(cardcounts, []int{1, 4}):
		hand = FOUROFAKIND
	case equal(cardcounts, []int{2, 3}):
		hand = FULLHOUSE
	case equal(cardcounts, []int{1, 1, 3}):
		hand = THREEOFAKIND
	case equal(cardcounts, []int{1, 2, 2}):
		hand = TWOPAIR
	case equal(cardcounts, []int{1, 1, 1, 2}):
		hand = ONEPAIR
	case equal(cardcounts, []int{1, 1, 1, 1, 1}):
		hand = HIGHCARD
	default:
		panic(fmt.Sprintf("Unexpected card counts: %v", cardcounts))
	}

	return &Hand{hand, cards, bid}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

type Game [](*Hand)

func (g Game) Len() int      { return len(g) }
func (g Game) Swap(i, j int) { g[i], g[j] = g[j], g[i] }
func (g Game) Less(i, j int) bool {
	a := *g[i]
	b := *g[j]

	if a.Type != b.Type {
		return a.Type < b.Type
	}

	cardValues := map[rune]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	for k := 0; k < 5; k++ {
		if cardValues[a.Cards[k]] != cardValues[b.Cards[k]] {
			return cardValues[a.Cards[k]] < cardValues[b.Cards[k]]
		}
	}

	return false
}
