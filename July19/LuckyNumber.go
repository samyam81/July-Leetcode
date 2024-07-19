package main

func luckyNumbers (matrix [][]int) []int {
    minArr := []int{}
	for i := 0; i < len(matrix); i++ {
		minNo := matrix[i][0]
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < minNo {
				minNo = matrix[i][j]
			}
		}
		minArr = append(minArr, minNo)
	}
	maxArr := []int{}
	for i := 0; i < len(matrix[0]); i++ {
		maxEle := 0
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] > maxEle {
				maxEle = matrix[j][i]
			}
		}
        maxArr = append(maxArr, maxEle)
	}
	res:=[]int{}
	for i:=0;i<len(minArr);i++ {
		for j:=0;j<len(maxArr);j++{
            if minArr[i]==maxArr[j]{
                res=append(res,minArr[i])
            }
        }
	}
	return res
}