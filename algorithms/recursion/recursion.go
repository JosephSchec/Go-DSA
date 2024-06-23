package recursion

import "fmt"

type Point struct {
	X int
	Y int
}

/* MAZE looks like this
   [
	"########E###"
	"##        ##"
	"##S#########"
   ]
*/

/* Base cases
* Off the map
* Hit a wall
* At the end
* Have we seen it
 */

var directions = [][]int{
	{-1, 0}, //up
	{1, 0},  //right
	{0, -1}, //down
	{0, 1},  //left
}

func Walk(maze []string, wall string, curr Point, end Point, seen [][]bool, paths *[]Point) bool {
	// off the map
	if curr.X < 0 || curr.X >= len(maze[0]) || curr.Y < 0 || curr.Y >= len(maze[0]) {
		return false
	}
	// hit a wall
	if string(maze[curr.Y][curr.X]) == wall {
		return false
	}
	// at the end
	if curr.X == end.X && curr.Y == end.Y {
		*paths = append(*paths, end)
		return true
	}
	if seen[curr.Y][curr.X] {
		return false
	}
	// pre
	seen[curr.Y][curr.X] = true
	*paths = append(*paths, Point{X: curr.X, Y: curr.Y})

	//recurse
	for i := 0; i < len(directions); i++ {
		newCurr := directions[i]
		if (Walk(maze, wall, Point{X: curr.X + newCurr[0], Y: curr.Y + newCurr[1]}, end, seen, paths)) {
			return true
		}
	}

	//post
	// in the scenario where there is a wall in the way we want to go to the last point and try again
	// this is why it is important to have a `seen` attribute since when we go back we don't want to try this direction again
	// rather we go to the next direction
	length := len(*paths)
	s := *paths
	s = s[:length-1]

	*paths = s

	return false
}
func SolveMaze(maze []string, wall string, start Point, end Point) ([]Point, error) {
	var paths []Point
	var seen [][]bool
	for i := 0; i < len(maze); i++ {
		seen = append(seen, make([]bool, len(maze[i])))
	}
	Walk(maze, wall, start, end, seen, &paths)
	return paths, nil
}

func RunSampleMaze() {
	maze := []string{
		"######E####",
		"###      ##",
		"##    #####",
		"###S#######",
	}
	peekValue, err := SolveMaze(maze, "#", Point{X: 3, Y: 2}, Point{X: 6, Y: 0})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Values", peekValue)
	}
}
