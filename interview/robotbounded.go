package main

import "fmt"

func main() {
	isRobotBounded("GGLLGG")
}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dirX, dirY := 0, 1
	for k := 0; k < 4; k++ {

		for i := 0; i < len(instructions); i++ {
			if instructions[i] == 'G' {
				x, y = x+dirX, y+dirY
			} else if instructions[i] == 'L' {
				dirX, dirY = dirY*-1, dirX
			} else if instructions[i] == 'R' {
				dirX, dirY = dirY, dirX*-1
			}

		}
	}
	fmt.Println(x, y, dirX, dirY)
	return (x == 0 && y == 0) && (dirX == 0 && dirY == 1)
}

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	dirX, dirY := 0, 1
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'G' {
			x, y = x+dirX, y+dirY
		} else if instructions[i] == 'L' {
			dirX, dirY = dirY*-1, dirX
		} else if instructions[i] == 'R' {
			dirX, dirY = dirY, dirX*-1
		}
	}
	return (x == 0 && y == 0) || (dirX != 0 || dirY != 1)
}
