# 2023 Day 22: Sand Slabs

https://adventofcode.com/2023/day/22

## Part 1

Given bricks which are rectangular parallelepipeds (3D rectangles) having integer sizes and integer coordinates.
The task is to "fall" them onto the ground and then find out number of bricks which could be safely removed, i.e. if you remove this brick,
not a single brick will further fall.
Falling works like in Tetris: it stops if some part of the brick cannot move down because of other brick (or ground).

### Solution

Outline:
1. Sort bricks by bottom Z coordinate in non-decreasing order.
2. Land each brick on the surface below starting from the lowest one, memorising which other brick(s) each brick stands on.
3. Count bricks which are the only base for the other brick - we cannot remove them, they are "supporting".
4. Return number of bricks minus count on the previous step.

#### Technical note 1

For efficient landing of the brick a "top view" data stracture is used, for each (x, y) store two numbers:
* index of this uppermost brick at this point
* z coordinate of the top of this brick

If you need to land a brick with coordinates (x0..x1, y0..y1), you just find maximum z for that 2D rectangle in the "top view", and this will be
the basis for new position of the brick after falling.

#### Complexity

Time complexity:
* `O(N*log(N))` for sort, N = number of bricks
* `O(N * S)` for falling, S = max square of the brick in XY projection
* `O(N)` for search of "supporting" bricks

Total: `O(N*log(N)+N*S)`

#### Alternative: DFS from the ground

There is an alternative for step 3. Note that after fall was done not a single brick is flying in the air, i.e. they are all standing on some other bricks
(or on the ground). Which means that if we will "jump" from the ground up to the next brick which is lying on ground and so on, we can visit all the bricks.
Then we can remove one brick and see whether all the other bricks are reachable from the ground. 

The compexity would be bigger cause DFS works for `O(U+V)` = `O(N*N)` in worst case, but input test data a not so big and it works really fast.

#### Alternative 2: Simulation

In my test file all the bricks are contained in a 3D rectangle 10x10x308, so you can just emulate space for 30K cells.
This would work, a bit slower, and I am not sure it is faster to implement.


## Part 2

Now instead of finding non-supporting bricks we now need to find out how many brick would fall if one of the non-supporting brick was removed.
Then return sum of all counts.

Solution is based on DFS from the ground:
* remove i-th brick
* start dfs from the ground
* brisk that have not been visited are falling down

### Complexity

`O(N*N)` in worst case, still really fast for our test data.
