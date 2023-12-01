# Advent of Code 2023

[![Tests](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml)
[![Stars: 2](https://img.shields.io/badge/⭐_Stars-2-yellow)](https://adventofcode.com/2023)

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
