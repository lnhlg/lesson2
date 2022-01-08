package lesson2

//NumSubmatrixSumTarget: 计算元素和为目标值的子矩阵数量
/*parameter
matrix: 整型矩阵
target: 目标值
return: 符合条件的矩阵数量
*/
func NumSubmatrixSumTarget(matrix [][]int, target int) (ans int) {

	m, n := len(matrix), len(matrix[0])
	if m < n {

		ans = upDownSum(matrix, m, n, target)
	} else {

		ans = leftRightSum(matrix, m, n, target)
	}

	return
}

//upDownSum: 以上下边界的形式计算
/*parameter
matrix: 整型矩阵
m, n: 矩阵的长宽
target: 目标值
return: 符合条件的矩阵数量
*/
func upDownSum(matrix [][]int, m, n, target int) (ans int) {

	for upRow := 0; upRow < m; upRow++ {

		rowSums := make([]int, n)
		for downRow := upRow; downRow < m; downRow++ {

			for col := 0; col < n; col++ {

				rowSums[col] += matrix[downRow][col]
			}

			ans += SubarraySum(rowSums, target)
		}
	}

	return
}

//leftRightSum: 以左右边界的形式计算
/*parameter
matrix: 整型矩阵
m, n: 矩阵的长宽
target: 目标值
return: 符合条件的矩阵数量
*/
func leftRightSum(matrix [][]int, m, n, target int) (ans int) {

	for leftCol := 0; leftCol < n; leftCol++ {

		colSums := make([]int, m)
		for rightCol := leftCol; rightCol < n; rightCol++ {

			for row := 0; row < m; row++ {

				colSums[row] += matrix[row][rightCol]
			}

			ans += SubarraySum(colSums, target)
		}
	}

	return
}
