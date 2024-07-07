package main

func numWaterBottles(numBottles int, numExchange int) int {
	totalBottles := numBottles

	for numBottles >= numExchange {
		totalBottles += numBottles / numExchange
		numBottles = (numBottles / numExchange) + (numBottles % numExchange)
	}

	return totalBottles
}