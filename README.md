# Advent of Code 2023

[![Tests](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml)
[![Stars: 8](https://img.shields.io/badge/⭐_Stars-8-yellow)](https://adventofcode.com/2023)

## Plan for This Year

This year I am going to try a couple of new things. First, I am going to try to
use the delve go debugger rather than put in print statements while I debug. My
hope is that by the end of December I will be much more familiar with the
debugger. Second, I want to try asking some generative AIs for helpful functions
to see how it improves my speed. I was considering Github copilot, but I just
can't give up my current editor, [helix](https://helix-editor.com/), to use
vscode, and I don't really want to go down that neovim plugin rabbit hole
anymore. 

I may use codespaces a bit. I've added some permissions so that I can clone my
private inputs submodule just in case, but it's hard to beat the setup I use to
write code every day.

This year I also created a new Advent of Code template in the
[devries/aoc_template](https://github.com/devries/aoc_template) repository. It
compiles everything into a single executable and times how long it takes to run
each problem, as well as generates a template for each day. It was also an
opportunity to experiment a bit with code generation.

## Solutions

- [Day 1: Trebuchet?!](https://adventofcode.com/2023/day/1) - [part 1](day01p1/solution.go), [part 2](day01p2/solution.go)

  Initially for the second part I put together a regular expression to search
  for the spelled out digits, however it is possible to have overlap between 
  these words, for example "twone", "sevenine", or "threeight." Regular
  expressions do not capture overlapping matches, so I switched to cycling 
  through the names and finding the indecies of each word or digit (if any)
  using `strings.Index`. Once you find a match you have to look for more
  matches starting at the next character after your previous search and I did
  have some indexing issues which I had to debug. I later realized I could use
  `strings.LastIndex` to make this much easier.

- [Day 2: Cube Conundrum ](https://adventofcode.com/2023/day/2) - [part 1](day02p1/solution.go), [part 2](day02p2/solution.go)

  Most of this problem was parsing. I just did a lot of splitting on substrings.
  First I split on ": " to separate the game id from the draws, then I split on
  "; " to separate the individual draws, then on ", " to split to the colors.
  After that it was straightforward. I missed an opportunity to use the debugger
  to check my parsing and instead put in a `Println`. 

- [Day 3: Gear Ratios ](https://adventofcode.com/2023/day/3) - [part 1](day03p1/solution.go), [part 2](day03p2/solution.go)

  The key here was just figuring out how to store the schematic data so it would
  be useful for answering the questions. During parsing I recorded the positions
  of the symbols in a map indexed by position. I recorded the numbers as an
  array of structs including the value and the start and end position. Then I
  could interate over all the numbers, find the surrounding points, and see what
  symbols were around them. I created an array of numbers for each gear object
  I found and then could iterate through those arrays to find gears with 
  exactly two adjacent numbers. 

- [Day 4: Scratchcards](https://adventofcode.com/2023/day/4) - [part 1](day04p1/solution.go), [part 2](day04p2/solution.go)

  The challenge for today was competing without any power. Eventually I set up a
  hotspot on my phone to connect via a very slow (1 bar LTE) connection and
  read the problem as well as download the input. I am so glad that the Advent
  of Code site is all text. As for the problem? I am cold and tired so I am not
  sure it was my best code, but I just iterated through the cards while creating
  an array of the number of copies of the card which I called `multiplier`.
  Because go arrays default to 0, I decided it would be easier if the multiplier
  array just had the count minus 1 so that a multiplier of 0 would mean 1 copy.
  Everything seems to have worked out in the end, though I still don't have
  power, and it's not getting any warmer.
