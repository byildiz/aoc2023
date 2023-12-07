# My AoC 2023 Solutions in Go
I am learning Go and want to improve it. One of the best things to do when learning a new language is to write code in that language. Because of that, I am solving Advent of Code 2023 (https://adventofcode.com/2023). Besides, it improves my problem-solving ability. I am sharing my solutions both for myself and for those of you who want to improve yourselves. Good luck with your journey!

## Notes
### Day 5
The input numbers are too big for simple implementation. We don't need to handle each seed number individually instead we can think them as a consecutive sets and finding intersections between one of those sets and on map input can be done at O(1). The rest is implementation details. My solution assumes there is no intersection between the found intersections of the sources and the maps.

