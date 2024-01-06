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
For the first part, create every possible pattern and check if it fits the template. For the second part, dynamic programming is required to make the solution efficient enough. Otherwise, trying every possible pattern would not be feasible in terms of time. For second part I got the idea from [Jonathan's solution](https://www.youtube.com/watch?v=xTGkP2GNmbQ).

### Day 13
We can think of each row and column as a binary number (e.g., `#.## = 1011`). For the first part, we can search for mirror positions using brute force. For the second part, observe that if there is only one flipped bit, then the XOR of the number with the flipped bit and the original number is a power of 2. We can determine if a number is a power of 2 [efficiently](https://stackoverflow.com/a/600306/1306183). Using this observation, we can find a mirror point that contains only one flipped bit.

### Day 14
The key observation is that after a certain number of cycles of tilting, the state of the platform begins to repeat itself. At that point, we can stop tilting and calculate the remaining number of cycles from the 1 billion cycles. After determining that number, we can perform those additional cycles of tilting. Lastly, we can calculate the final load.

### Day 15
We can maintain a hash map for each box. When inserting a new lens into a box, we need to keep track of the lens order. My approach is to use a counter for each box. By inserting a new lens with the current value of the corresponding counter, we can then sort them afterward.

### Day 16
We can use any tree search algorithm here. To detect and break potential infinite loops, we can maintain a dynamic programming (DP) lookup table. The state for the DP lookup table includes the current location and the direction from which the beam enters the tile.

### Day 17
We can use any tree search algorithm with a lookup table. The state will consist of position, direction, and repetition. The tricky point in part 2 is that we have to enter the last block at least 4 repetitions.

### Day 18
The total volume of the lagoon is the sum of the perimeter of the loop and the number of points within the loop. For the first part, I created a matrix, drew the loop on it, and counted the points inside the loop using the DFS algorithm. However, this approach is not efficient enough for the second part of the question. I couldn't devise a solution for the second part until I understood [Pick's theorem](https://en.wikipedia.org/wiki/Pick%27s_theorem), [Green's theorem](https://en.wikipedia.org/wiki/Green%27s_theorem), and the [Shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula). For more information, please refer to [Jonathan's explanation](https://www.youtube.com/watch?v=UNimgm_ogrw).

### Day 19
In the first part of the question, we can traverse the tree for each part and determine whether the part is accepted or not. For the second part, trying all the parts would be too inefficient. Instead, we can identify all the paths that end with A and determine the boundaries (minimum and maximum accepted values for properties x, m, a, and s) for each property along each path. The number of parts accepted by a path is calculated by multiplying the difference between the maximum and minimum values for each property. Note that the result of this subtraction can be negative, so we should disregard those values.

### Day 20
We can view the Flip-Flop module as a JK flip-flop and the Conjunction as a NAND gate. The first part of the problem can be solved by simulating the circuit. For the second part, we need to investigate the circuit to gain insight into the internal wiring of `rx`. `rx` is the output of a conjunction named `df`, which has four inputs. `rx` receives a low pulse only when all inputs of `df` become high simultaneously. The inputs of `df` exhibit a cyclic behavior with varying periods. Therefore, the answer will be the least common multiple of these periods. For more information, please refer to [Jonathan's explanation](https://www.youtube.com/watch?v=3STpz-M-wiw).

### Day 21
The first part of the problem can be solved with breadth-first search (BFS). For the second part, I had to understand [Jonathan's solution](https://www.youtube.com/watch?v=C2dmxCGGH1s). Afterward, I developed my solution. Please refer to it for more information. The main observation for the solution is that after a certain number of tiles, the distances between the two points at the same location in consecutive tiles are equal to the number of rows (this is true if the number of rows is equal to the number of columns).

### Day 22
My approach to this problem is as follows: First, starting from the bottom and moving upwards, I relocate each block to its farthest possible position in the -z direction. Secondly, I calculate a dependency matrix that indicates which block supports another. With this matrix, I can determine how many supporters each block has. The steps taken up to this point apply to both the first and second parts of the problem. At this juncture, my approach diverges depending on the part of the problem.

For the first part, a block can be safely disintegrated only if all the blocks it supports have two or more supporters.

For the second part, if a block is disintegrated, I can safely reduce the number of supporters for the blocks it originally supported by one. If, after this reduction, any of these blocks have no supporters, it indicates that the block will fall. I can then recursively apply the same logic to those blocks and count all the blocks that end up with no support. By applying this procedure to each block and summing up the blocks with no supporters, we can determine the answer.

### Day 23
The first part of the problem can be solved by exploring all possible paths using depth-first search (DFS), given that the paths form a directed graph. In the second part of the problem, since there are no directions on the edges, the problem transforms into the traveling salesman problem, which is NP-Hard. Brute-forcing with DFS on the raw input is not efficient. To solve the problem within a reasonable time frame, we need to condense the map into a more compact graph. This compact graph can be created as follows: Vertices will represent the entry point, destination point, and all branch points. Edges will connect these vertices, with weights assigned based on the length of the path between them.

### Day 24
For part one, we need to find the intersection point for each pair of vectors (the intersection point can be determined using [line-line intersection](https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection) with two points from each vector) if it exists. Then, we count the pairs where an intersection point exists, and this intersection point lies in the direction of the vectors.

For part two, I got the idea from a comment below [Jonathan's solution](https://www.youtube.com/watch?v=vZa2jErpSg8). Here is my explanation for it:
Let there are two hailstones (h1  and h2) with `x1`, `x2` starting positions and `vx1` and `vx2` velocities in x direction. The rock's starting position is `x` with `vx` speed in x direction. Collision times with the rock are `t1` and `t2` for h1 and h2. The two basic equations we have are:
```
x + vx*t1 = x1 + vx1*t1
x + vx*t2 = x2 + vx2*t2
```
When we subtract two equations each other, we get `vx(t1-t2) = x1 + vx1*t1 - x2 - vx2*t2`. If `vx1` and `vx2` are equal, the equation becomes `vx(t1-t2) = x1 - x2 + vx1(t1 - t2)`. We can rearrange it: `(t1-t2) = (x1 - x2) / (vx - vx1)`. If we assume `t1` and `t2` are both integers, we get `(x1 - x2) % (vx - vx1) = 0`.
Starting from a set of velocities for each axis, we can eleminate some of them for each hailstone pair with equal velocities in one of the x, y an z axis. We can find the correct velocities by searching in the remaining sets.

Let's consider two hailstones `h1` and `h2` with starting positions `x1` and `x2`, and velocities `vx1` and `vx2` in the x-direction, respectively. The rock has an initial position `x` and a speed `vx` in the x-direction. The collision times for `h1` and `h2` with the rock are `t1` and `t2`, respectively. The two fundamental equations are:
```
x + vx*t1 = x1 + vx1*t1
x + vx*t2 = x2 + vx2*t2
```
When we subtract the second equation from the first, we obtain `vx(t1-t2) = x1 + vx1*t1 - x2 - vx2*t2`. If `vx1` and `vx2` are equal, this simplifies to `vx(t1-t2) = x1 - x2 + vx1(t1 - t2)`. Rearranging this gives `(t1-t2) = (x1 - x2) / (vx - vx1)`. Assuming `t1` and `t2` are both integers, we have `(x1 - x2) % (vx - vx1) = 0`.

By starting with a set of velocities for each axis, we can eliminate some velocities for each hailstone pair that have equal velocities in one of the x, y, or z axes. The correct velocities can then be determined by searching within the remaining sets.

### Day 25
I wrote a [Python code](./day25/draw.py) to draw the graph. In the graph, the three edges are clearly visible. So, I manually removed them, and using a simple depth-first search, I calculated the answer.
