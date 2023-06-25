package leetcode

func spiralOrder(matrix [][]int) []int {
	end := 200
	i, j := 0, 0
	res := []int{}
	n, m := len(matrix), len(matrix[0])
	for len(res) < m*n {
		// 向左遍历
		for ; j < m && i < n && matrix[i][j] != end; j++ {
			res = append(res, matrix[i][j])
			matrix[i][j] = end
		}
		j--
		i++
		for ; i < n && j >= 0 && matrix[i][j] != end; i++ {
			res = append(res, matrix[i][j])
			matrix[i][j] = end
		}
		i--
		j--
		for ; j >= 0 && i >= 0 && matrix[i][j] != end; j-- {
			res = append(res, matrix[i][j])
			matrix[i][j] = end
		}
		j++
		i--
		for ; i >= 0 && j < m && matrix[i][j] != end; i-- {
			res = append(res, matrix[i][j])
			matrix[i][j] = end
		}
		i++
		j++
	}
	return res
}

const (
	up    = 1
	down  = 2
	left  = 3
	right = 4
)

func move(i, j, dict int) (int, int) {
	x, y := getXY(dict)
	return i + x, j + y
}

func moveByStep(i, j, dict, step int) (int, int) {
	x, y := getXY(dict)
	return i + x*step, j + y*step
}

var dictMap = map[int][]int{
	up:    []int{-1, 0},
	down:  []int{1, 0},
	left:  []int{0, -1},
	right: []int{0, 1},
}

func getXY(dict int) (int, int) {
	return dictMap[dict][0], dictMap[dict][1]
}
