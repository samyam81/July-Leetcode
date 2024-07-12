package main

func maximumGain(s string, x int, y int) int {
    ans, curr := 0, 0

    if x >= y {
        curr = 0
        curr, s = helper(s, "ab", x)
        ans += curr
        curr, _ = helper(s, "ba", y)
        ans += curr
    } else {
        curr = 0
        curr, s = helper(s, "ba", y)
        ans += curr
        curr, _ = helper(s, "ab", x)
        ans += curr
    }
    return ans
}

func helper(s string, sub string, x int) (int, string) {
    var score int
    var stack []byte
    for i := 0; i < len(s); i++ {
    if s[i] == sub[1] && len(stack) > 0 && stack[len(stack)-1] == sub[0] {
        stack = stack[:len(stack)-1] 
        score += x
    } else {
        stack = append(stack, s[i])
    }
    }
    return score, string(stack)
}