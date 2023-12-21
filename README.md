# My AoC 2023 Solutions in Go
I am learning Go and want to improve it. One of the best things to do when learning a new language is to write code in that language. Because of that, I am solving Advent of Code 2023 (https://adventofcode.com/2023). Besides, it improves my problem-solving ability. I am sharing my solutions both for myself and for those of you who want to improve yourselves. Good luck with your journey!

## Notes

### Day 5
The input numbers are too big for simple implementation. We don't need to handle each seed number individually instead we can think them as a consecutive sets and finding intersections between one of those sets and on map input can be done at O(1). The rest is implementation details. My solution assumes there is no intersection between the found intersections of the sources and the maps.

### Day 6
Part 1 can be solved by brute force; however, for part 2, binary search is needed since the distance is too big for a brute force search.

### Day 7
If we pretend a hand as a hexadecimal number (A: E, K: D, Q: C, J: B, T: A, 9: 9, ...), then a card will be represented in 4 bits. The total value of a hand can fit into 20 bits. The type of a hand can be determined by counting the characters and sorting those counts. The strength of a hand is its value plus 20 bits left-shifted by its type.

In part 2, we can substitute Js with 1s for value calculation. To find the type of a hand, we should convert jokers to the most frequent card in order to make the hand as strong as possible.

### Day 8
For the second part, we can first find the required number of steps to cycle each of the XXA->XXZ paths, and then the answer will be least common multiple of those steps. The solution assumes that we are right at the beginning of each of the paths.

### Day 9
```
p
  p1
a       p2
  (b-a)          p3
b       (c-2b+a)             p4
  (c-b)          (d-3c+3b-a)                p5
c       (d-2c+b)             (e-4d+6c-4b+a)                     0
  (d-c)          (e-3d+3c-b)                (f-5e+10d-10c+5b-a)   0
d       (e-2d+c)             (f-4e+6d-4c+b)                     0
  (e-d)          (f-3e+3d-c)                n5
e       (f-2e+d)             n4
  (f-e)          n3
f       n2
  n1
n
```

**Observations**
1. The coefficients of the numbers are [the binomial coefficients](https://en.m.wikipedia.org/wiki/Binomial_coefficient).
1. `n = f + n1 = f + (f-e) + n2 = f + (f-e) + (f-2e+d) + n3 = ...`
1. `p = a - p1 = a - (b-a) + p2 = a - (b-a) + (c-2b+a) - p3 = ...`

The rest is the implementation details.

### Day 10
For part 2, my approach is first expand the map (specifically double both rows and cols and fill them first with dots and then connect main loop with | and -) and then find the main loop. After finding main loop, starting from edge tiles, mark all the tiles outside of the main loop (tart from a tile and if it is not part of main loop mark it and recursively look around it to mark them, too). Once all the marking is done, we can delete the added rows and columns and savely count the unmarked ones as inner tiles.

#### Example:
**Original Map:**
```
..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........

```

**Expanded Map:**
```
....................
....................
..S-------------7...
..|.............|...
..|.F---------7.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.|.........|.|...
..|.L---7.F---J.|...
..|.....|.|.....|...
..|.....|.|.....|...
..|.....|.|.....|...
..L-----J.L-----J...
....................
....................
....................
```

**Main Loop: 2**
```
00000000000000000000
00000000000000000000
00222222222222222000
00200000000000002000
00202222222222202000
00202000000000202000
00202000000000202000
00202000000000202000
00202000000000202000
00202000000000202000
00202222202222202000
00200000202000002000
00200000202000002000
00200000202000002000
00222222202222222000
00000000000000000000
00000000000000000000
00000000000000000000
```

**Main Loop: 2, Outer Tiles: 1**
```
11111111111111111111
11111111111111111111
11222222222222222111
11200000000000002111
11202222222222202111
11202111111111202111
11202111111111202111
11202111111111202111
11202111111111202111
11202111111111202111
11202222212222202111
11200000212000002111
11200000212000002111
11200000212000002111
11222222212222222111
11111111111111111111
11111111111111111111
11111111111111111111
```

**Shrunken Map:**
```
1111111111
1222222221
1222222221
1221111221
1221111221
1222222221
1200220021
1222222221
1111111111
```
### Day 11
Empty row means that we should add (expandCoeff-1) to the row numbers of each galaxy which has greater row number than the empty row. The same goes for empty columns. We can do expantion in an efficient way: Sort both the galaxies by their row numbers and empty rows ids ascending then start expantion from reverse because the rows effects the rows coming after them. For empty columns, sort both galaxies by their column numbers and empty column ids again asceding and do the expantion again reverse for the same reason.

### Day 12
For the first part, create every possible pattern and check if it fits the template. For the second part, dynamic programming is required to make the solution efficient enough. Otherwise, trying every possible pattern would not be feasible in terms of time.

### Day 13
We can think of each row and column as a binary number (e.g., `#.## = 1011`). For the first part, we can search for mirror positions using brute force. For the second part, observe that if there is only one flipped bit, then the XOR of the number with the flipped bit and the original number is a power of 2. We can determine if a number is a power of 2 [efficiently](https://stackoverflow.com/a/600306/1306183). Using this observation, we can find a mirror point that contains only one flipped bit.

### Day 14
The key observation is that after a certain number of cycles of tilting, the state of the platform begins to repeat itself. At that point, we can stop tilting and calculate the remaining number of cycles from the 1 billion cycles. After determining that number, we can perform those additional cycles of tilting. Lastly, we can calculate the final load.

### Day 15
We can maintain a hash map for each box. When inserting a new lens into a box, we need to keep track of the lens order. My approach is to use a counter for each box. By inserting a new lens with the current value of the corresponding counter, we can then sort them afterward.

### Day 16
We can use any tree search algorithm here. To detect and break potential infinite loops, we can maintain a dynamic programming (DP) table. The state for the DP table includes the current location and the direction from which the beam enters the tile.
