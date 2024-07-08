package main

func findTheWinner(n int, k int) int {
    i:=1
    for j:=2;j<=n;j++{
        i+=k
        i=(i-1)%j+1
    }
    return i
}