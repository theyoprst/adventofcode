## Making AI to solve problems

It appears that at least simple AOC tasks (or even maybe all of them) can be solved with good-enough-in-coding LLM.
E.g.
- Chat GPT - tested in free account, good enough
- Claude - didn't test
- Meta-Llama-3.1-405b - tested with Nebius AI Studio, good enough

The prompt:

```
Solve an advent of code problem in Go. The problem has two parts, the first part will be in the first user message, the second one will follow in the next user message. Your solution must be a function with signature `SolvePart{PART_NUMBER}(lines []string) any`. So no error is returned.  Do not validate input, assume that it is already valid.
Do not write imports, main function, usage examples or any other additional text.
To split items by spaces prefer using strings.Fields(). Refrain from using variable names with indices like list1, list2, etc.
You can use the following helpers:

func must.Atoi(s string) int
    must.Atoi is a wrapper around strconv.Atoi that panics on error.

func aoc.Abs[T constraints.Signed | constraints.Float](a T) T
    aoc.Abs returns absolute value of `a`.
```
