package main

func averageWaitingTime(customers [][]int) float64 {
    totalWaitingTime := 0
    currentTime := 0
    
    for _, customer := range customers {
        arrival := customer[0]
        service := customer[1]
        
        if currentTime < arrival {
            currentTime = arrival
        }
        
        waitingTime := currentTime - arrival
        totalWaitingTime += waitingTime + service
        currentTime += service
    }
    
    return float64(totalWaitingTime) / float64(len(customers))
}
