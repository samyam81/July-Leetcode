package main

func threeConsecutiveOdds(arr []int) bool {
    if len(arr)<=2 {
		return false
	}
	var count uint=0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2==1 {
			count++
			if count==3 {
				return true
			}
		} else{
			count=0
		}
	}
	return false
}