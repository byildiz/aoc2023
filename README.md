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

