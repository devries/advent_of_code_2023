# Advent of Code 2023

[![Tests](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml/badge.svg)](https://github.com/devries/advent_of_code_2023/actions/workflows/main.yml)
[![Stars: 30](https://img.shields.io/badge/⭐_Stars-30-yellow)](https://adventofcode.com/2023)

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
| 10  | 1    | 59.225004ms  |
| 10  | 2    | 119.835531ms |
| 11  | 1    | 919.568362ms |
| 11  | 2    | 908.9021ms   |
| 12  | 1    | 7.49596ms    |
| 12  | 2    | 129.23797ms  |
| 13  | 1    | 708.271µs    |
| 13  | 2    | 711.826µs    |
| 14  | 1    | 3.185616ms   |
| 14  | 2    | 1.734193401s |
| 15  | 1    | 460.069µs    |
| 15  | 2    | 3.105932ms   |

## Solutions

- [Day 1: Trebuchet?!](https://adventofcode.com/2023/day/1) - [⭐ part 1](day01p1/solution.go), [⭐ part 2](day01p2/solution.go)

  Initially for the second part I put together a regular expression to search
  for the spelled out digits, however it is possible to have overlap between 
  these words, for example "twone", "sevenine", or "threeight." Regular
  expressions do not capture overlapping matches, so I switched to cycling 
  through the names and finding the indecies of each word or digit (if any)
  using `strings.Index`. Once you find a match you have to look for more
  matches starting at the next character after your previous search and I did
  have some indexing issues which I had to debug. I later realized I could use
  `strings.LastIndex` to make this much easier.

- [Day 2: Cube Conundrum ](https://adventofcode.com/2023/day/2) - [⭐ part 1](day02p1/solution.go), [⭐ part 2](day02p2/solution.go)

  Most of this problem was parsing. I just did a lot of splitting on substrings.
  First I split on ": " to separate the game id from the draws, then I split on
  "; " to separate the individual draws, then on ", " to split to the colors.
  After that it was straightforward. I missed an opportunity to use the debugger
  to check my parsing and instead put in a `Println`. 

- [Day 3: Gear Ratios ](https://adventofcode.com/2023/day/3) - [⭐ part 1](day03p1/solution.go), [⭐ part 2](day03p2/solution.go)

  The key here was just figuring out how to store the schematic data so it would
  be useful for answering the questions. During parsing I recorded the positions
  of the symbols in a map indexed by position. I recorded the numbers as an
  array of structs including the value and the start and end position. Then I
  could interate over all the numbers, find the surrounding points, and see what
  symbols were around them. I created an array of numbers for each gear object
  I found and then could iterate through those arrays to find gears with 
  exactly two adjacent numbers. 

- [Day 4: Scratchcards](https://adventofcode.com/2023/day/4) - [⭐ part 1](day04p1/solution.go), [⭐ part 2](day04p2/solution.go)

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

- [Day 5: If You Give A Seed A Fertilizer](https://adventofcode.com/2023/day/5) - [⭐ part 1](day05p1/solution.go), [⭐ part 2](day05p2/solution.go)

  This is the sort of problem where it is prohibitively large to calculate the
  conversion of every individual element in a range of integers, however the
  given ranges of integers need to be handled in different ways depending on
  where they exist within portions of those ranges. Rather then iterate through
  each range, the key is finding the subranges which are handled in the same way
  and calculate how that range as a whole will be modified. As you continue to
  do this the number of ranges grows, but it will always be far fewer
  calculations than tracking how each individual element is handled.

- [Day 6: Wait For It](https://adventofcode.com/2023/day/6) - [⭐ part 1](day06p1/solution.go), [⭐ part 2](day06p2/solution.go)

  This was just solving the quadratic equation for the times when the distance
  was equal to the time of the race. There was a bit of fiddling with checking
  if the distance was greater than or equal to the winning distance, but all in
  all that was the gist of it.

- [Day 7: Camel Cards](https://adventofcode.com/2023/day/7) - [⭐ part 1](day07p1/solution.go), [⭐ part 2](day07p2/solution.go)

  I thought this was an interesting one. Initially it involved parsing the card
  values and then scoring based on the count of cards in the hand and the values
  of the cards, but in the second part adding a wildcard put in a twist. I added
  the joker counts to the card with the highest count, so for example if I had
  two 3s, one J, and two other cards, I would add the J to the 3s making three
  of a kind. I got a bit hung up on the golang sorting pattern, maybe because I
  hadn't had my coffee?

- [Day 8: Haunted Wasteland](https://adventofcode.com/2023/day/8) - [⭐ part 1](day08p1/solution.go), [⭐ part 2](day08p2/solution.go)

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

- [Day 9: Mirage Maintenance](https://adventofcode.com/2023/day/9) - [⭐ part 1](day09p1/solution.go), [⭐ part 2](day09p2/solution.go)

  I thought about trying to somehow be clever and calculate only as much as I had
  to at the edges of the sequences, but then I thought each sequence is short so
  I just followed the procedure in the example. It turned out to be very
  straightforward.

- [Day 10: Pipe Maze](https://adventofcode.com/2023/day/10) - [⭐ part 1](day10p1/solution.go), [⭐ part 2](day10p2/solution.go)

  The first part was streightforward, follow the pipe in both directions until
  your search meets at the farthest point from the start. The second part took
  some time for me to think about. Eventually I settled on counting pipe
  intersections on a walk directly to the north from the point of interest. A
  point on the outside would intersect the pipe an even number of times, while
  a point on the outside would intersect an odd number of times.

- [Day 11: Cosmic Expansion](https://adventofcode.com/2023/day/11) - [⭐ part 1](day11p1/solution.go), [⭐ part 2](day11p2/solution.go)

  Straightforward again. I used the [github.com/devries/combs](https://github.com/devries/combs)
  library to get all the pairs of galaxies and then summed over the columns and
  rows that separate them, multiplying by the expansion factor if the column or
  row did not contain a galaxy.

- [Day 12: Hot Springs](https://adventofcode.com/2023/day/12) - [⭐ part 1](day12p1/solution.go), [⭐ part 2](day12p2/solution.go)

  This one was very tough for me. I spent a lot of time making iterators that would
  return the next potential valid row with one additional group filled in, but
  I wasn't able to memoize that method and I did not take into account the idea
  that there was a maximum number of working spring spaces I could put in before
  adding the next group. I ended up parameterizing solution counts in a memoizable
  state which, for each sequence, was the number of groups already accounted for
  and the starting position in the sequence. I would then run through all
  possible positions of the next group up to the total amount of buffer space
  I had available and find counts for those. I find these kinds of problems
  very difficult.

- [Day 13: Point of Incidence](https://adventofcode.com/2023/day/13) - [⭐ part 1](day13p1/solution.go), [⭐ part 2](day13p2/solution.go)

  I noticed all the maps were smaller than 64 in length and width, so I used a
  bitfield to store the locations of the rocks for every row and column. I then
  just had to check for symmetry from each row and column gap. For the second
  part I did the same comparison but checked to see if there was an off by 1
  bit issue between any two comparisons, and required one and only one of those
  to define a new symmetry axis. Unfortunately I tried to use subtraction rather
  than the XOR function, which of course caused a few errors. My extra test
  cases were made to track down those errors. I used Kernighan's bit counting
  algorithm to find the number of bits in the XOR difference. 

- [Day 14: Parabolic Reflector Dish](https://adventofcode.com/2023/day/14) - [⭐ part 1](day14p1/solution.go), [⭐ part 2](day14p2/solution.go)

  The tilting mechanic was fairly straighforward, though for the spin cycle I
  decided to write four separate loops rather than a more generic loop which
  would work for each direction. Obviously iterating 1,000,000,000 times is not
  practical, however the arrangement of rocks should begin cycling through a
  sequence of repeated positions, so we just need to find where that cycle
  starts and ends, and then jump ahead N cycles to just before the iterations
  are complete. The only problem is storing the map state. This tripped me up a
  bit as I tried to find a unique integer that summarized the map, but that's
  the job of a hash function. I took the sorted list of positions of rolling
  rocks, turned the coordinate into an integer, *sorted them*, and then wrote
  the binary encoding of those integers to a fnv-1a hash function. I then was
  able to look for when the hashes started to repeat. This is the first day
  that takes more than 1 second to run on the Raspberry Pi (though day 11 came
  close).

  So far I have not used any AI generated subroutines in my code. Today may have
  been a good day to try. I could have asked something like "write a function in
  go that calculates the hash of a slice of int32s." Having just asked, bing
  generated the following code:

    ```go
    import (
        "crypto/sha256"
        "encoding/binary"
    )

    func HashInt32Slice(s []int32) [32]byte {
        h := sha256.New()
        for _, i := range s {
            b := make([]byte, 4)
            binary.LittleEndian.PutUint32(b, uint32(i))
            h.Write(b)
        }
        return h.Sum(nil)
    }
    ```
    
  While programming it seems like I don't really want to shift gears to ask AI
  for help.

- [Day 15: Lens Library](https://adventofcode.com/2023/day/15) - [⭐ part 1](day15p1/solution.go), [⭐ part 2](day15p2/solution.go)

  This was also a very straightforward problem, although I did have to fix a
  number of bugs I introduced. This problem illustrates how a hash map works
  and then requires operations on the hashmapped values. I created a Lens object
  which I passed around by value, and forgot when I changed an attribute in that
  object, I need to reassign it to the lens in the box array.
