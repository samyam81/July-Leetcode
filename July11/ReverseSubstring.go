package main

func reverseParentheses(s string) string {
    stack := make([]int, 0)
    chars := []byte(s)
    
    for i := 0; i < len(chars); i++ {
        if chars[i] == '(' {
            stack = append(stack, i)
        } else if chars[i] == ')' {
            start := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            reverse(chars, start+1, i-1)
        }
    }
    result := make([]byte, 0, len(chars))
    for i := 0; i < len(chars); i++ {
        if chars[i] != '(' && chars[i] != ')' {
            result = append(result, chars[i])
        }
    }
    
    return string(result)
}

func reverse(chars []byte, left, right int) {
    for left < right {
        chars[left], chars[right] = chars[right], chars[left]
        left++
        right--
    }
}
