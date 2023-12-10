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

