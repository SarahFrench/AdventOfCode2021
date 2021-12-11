# Advent of Code 2021

## [Day 1](https://adventofcode.com/2021/day/1)

- Window structs are responsible for knowing if they need to report their latest sum & reset the window contents, and also responsible for ignoring numbers until their offset is satisfied
- WindowGroups contain multiple Windows and the constructor configures how big the Windows are in a group
- main function is responsible for feeding numbers to the WindowGroup, responding to any sums that are reported back, and tracking increases between sequential sums