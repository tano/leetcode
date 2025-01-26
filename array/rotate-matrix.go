package array

func rotateMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := len(matrix[i]) - 1; j > i; j-- {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	for i := 0; i < len(matrix); i++ {
		increasing := 0
		decreasing := len(matrix[i]) - 1
		for increasing <= decreasing {
			matrix[i][increasing], matrix[i][decreasing] = matrix[i][decreasing], matrix[i][increasing]
			increasing++
			decreasing--
		}
	}
}
