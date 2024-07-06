package main

func passThePillow(n int, time int) int {
    N := 2*n - 2
    x := time % N
    if x < n {
        return 1 + x
    } else {
        return 1 + (N - x)
    }
}
