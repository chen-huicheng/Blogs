package leetcode

var dir = []int{-1, 0, 1, 0, 0, -1, 0, 1}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				visit(grid, i, j)
			}
		}
	}
	return res
}
func visit(grid [][]byte, i, j int) {
	grid[i][j] = '0'
	for id := 0; id < 4; id++ {
		di, dj := i+dir[id], j+dir[id+4]
		if di < 0 || di >= len(grid) || dj < 0 || dj >= len(grid[i]) {
			continue
		}
		if grid[di][dj] == '1' {
			visit(grid, di, dj)
		}
	}
}
