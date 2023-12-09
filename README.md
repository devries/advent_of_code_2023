# Advent of Code 2023

[![Tests](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml)
[![Stars: 18](https://img.shields.io/badge/⭐_Stars-18-yellow)](https://adventofcode.com/2023)

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

## Efficiency

Although at times I have not found it, there should be a "[solution that completes
in at most 15 seconds on ten-year-old hardware.](https://adventofcode.com/2023/about)"

I am hoping this year to find solutions that will run in a reasonable amount of
time on a Raspberry Pi 4 Model B with 4 Gigabytes of RAM and a Cortex A72
processor. The Pi-4 is roughly [500-600 times faster](https://timeartisan.org/fftbench/)
than the first computer I purchased in the late 90s (200 MHz Pentium MMX with 64
Megabytes of RAM). Below I have the time it takes for the problem to finish on
the third run of my solution after compilation on my Raspberry Pi.

| Day | Part | Time Elapsed |
| :-- | :--- | :----------- |
| 1   | 1    | 558.698µs    |
| 1   | 2    | 3.230152ms   |
| 2   | 1    | 1.847944ms   |
| 2   | 2    | 1.950147ms   |
| 3   | 1    | 3.929274ms   |
| 3   | 2    | 4.738173ms   |
| 4   | 1    | 4.211679ms   |
| 4   | 2    | 3.712758ms   |
| 5   | 1    | 3.235819ms   |
| 5   | 2    | 1.154803ms   |
| 6   | 1    | 31.777µs     |
| 6   | 2    | 30.13µs      |
| 7   | 1    | 4.545508ms   |
| 7   | 2    | 5.002041ms   |
| 8   | 1    | 3.015281ms   |
| 8   | 2    | 23.987541ms  |
| 9   | 1    | 2.51753ms    |
| 9   | 2    | 2.292069ms   |

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

- [Day 5: If You Give A Seed A Fertilizer](https://adventofcode.com/2023/day/5) - [part 1](day05p1/solution.go), [part 2](day05p2/solution.go)

  This is the sort of problem where it is prohibitively large to calculate the
  conversion of every individual element in a range of integers, however the
  given ranges of integers need to be handled in different ways depending on
  where they exist within portions of those ranges. Rather then iterate through
  each range, the key is finding the subranges which are handled in the same way
  and calculate how that range as a whole will be modified. As you continue to
  do this the number of ranges grows, but it will always be far fewer
  calculations than tracking how each individual element is handled.

- [Day 6: Wait For It](https://adventofcode.com/2023/day/6) - [part 1](day06p1/solution.go), [part 2](day06p2/solution.go)

  This was just solving the quadratic equation for the times when the distance
  was equal to the time of the race. There was a bit of fiddling with checking
  if the distance was greater than or equal to the winning distance, but all in
  all that was the gist of it.

- [Day 7: Camel Cards](https://adventofcode.com/2023/day/7) - [part 1](day07p1/solution.go), [part 2](day07p2/solution.go)

  I thought this was an interesting one. Initially it involved parsing the card
  values and then scoring based on the count of cards in the hand and the values
  of the cards, but in the second part adding a wildcard put in a twist. I added
  the joker counts to the card with the highest count, so for example if I had
  two 3s, one J, and two other cards, I would add the J to the 3s making three
  of a kind. I got a bit hung up on the golang sorting pattern, maybe because I
  hadn't had my coffee?

- [Day 8: Haunted Wasteland](https://adventofcode.com/2023/day/8) - [part 1](day08p1/solution.go), [part 2](day08p2/solution.go)

  This is a series where you calculate the least common multiple of a set of 
  cycling states in order to find the interval over which all of the states
  fully run through their cycles. I initially printed out the step at which the
  cycle started and how many steps until it repeated. I noticed that the problem
  is contrived such that the number of steps it initially took to reach the
  desired end state was equal to the interval between times it hit that end
  state, which makes the problem a straighforward least common multiples problem.
  If the problem had not been contrived in that way, it is possible that there
  would not have been a period over which all starting states eventually
  synchronize so all ending states are reached at the same time.

- [Day 9: Mirage Maintenance](https://adventofcode.com/2023/day/9) - [part 1](day09p1/solution.go), [part 2](day09p2/solution.go)

  I thought about trying to somehow be clever and calculate only as much as I had
  to at the edges of the sequences, but then I thought each sequence is short so
  I just followed the procedure in the example. It turned out to be very
  straightforward.
