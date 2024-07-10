package main

func minOperations(logs []string) int {
    var depth int = 0

    for _, log := range logs {
        if log == "../" {
            if depth > 0 {
                depth--
            }
        } else if log != "./" {
            depth++
        }
    }

    return depth
}
