package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	out := make([][]int, len(rowSum))
	for row_sum_index, _ := range rowSum {
		out[row_sum_index] = make([]int, len(colSum))
		for col_sum_index, _ := range colSum {
			min_val := min(rowSum[row_sum_index], colSum[col_sum_index])
			rowSum[row_sum_index] -= min_val
			colSum[col_sum_index] -= min_val
			out[row_sum_index][col_sum_index] = min_val
		}
	}
	return out
}