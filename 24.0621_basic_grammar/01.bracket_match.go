package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    m := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }
    stack := make([]rune, 0)
    for _, v := range s {
        if (v == '[') || (v == '(') || (v == '{') {
            stack = append(stack, v)
        } else if (v == ']') || (v == ')') || (v == '}') {
            if len(stack) <= 0 {
                return false
            }
            lef_char := stack[len(stack)-1]
            rig_char := v
            if (m[lef_char] != rig_char) {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            fmt.Println("Invalid character!")
            return false
        }
    }
    return len(stack) == 0
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    text = strings.Trim(text, " \n")
    fmt.Println(text)
	fmt.Println(isValid(text))
}