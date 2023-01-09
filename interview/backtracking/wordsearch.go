package main

func main() {

}

func exist(board [][]byte, word string) bool {
	var found bool

	var dfs func(row, col int, index int)
	dfs = func(row, col int, index int) {
		if found {
			return
		}
		if index >= len(word) {
			found = true
			return
		}

		if row < 0 || col < 0 || row >= len(board) || col >= len(board[row]) {
			return
		}
		if board[row][col] != word[index] {
			return
		}

		tmp := board[row][col]
		board[row][col] = '#'

		direction := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
		for _, v := range direction {
			dfs(row+v.x, col+v.y, index+1)
		}
		board[row][col] = tmp
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs(i, j, 0)
			if found {
				return true
			}
		}
	}
	return false
}
