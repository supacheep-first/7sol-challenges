package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readInput() string {
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.ToUpper(strings.TrimSpace(input))
    if !strings.ContainsAny(input, "LR=") {
        fmt.Println("Invalid input format. Must contain 'L', 'R', or '='.")
        os.Exit(1)
    }
    return input
}

func isValid(last, next int, symbol byte) bool {
    return (symbol == 'L' && last > next) ||
        (symbol == 'R' && last < next) ||
        (symbol == '=' && last == next)
}

func findMinSequence(input string) string {
    minSum, best := len(input) * 10, ""
    var dfs func([]int)
    dfs = func(current []int) {

        if len(current) == len(input)+1 {
            sum := 0
            for _, v := range current {
                sum += v
            }
            if sum < minSum {
                minSum = sum
                best = strings.Trim(strings.Replace(fmt.Sprint(current), " ", "", -1), "[]")
            }
            return
        }
        last := current[len(current)-1]
        for next := 0; next <= len(input); next++ {
            if isValid(last, next, input[len(current)-1]) {
                dfs(append(current, next))
            }
        }
    }
    for start := 0; start <= 9; start++ {
        dfs([]int{start})
    }
    return best
}

func main() {
    input := readInput()
    fmt.Println("Input:", input)
    fmt.Println("Minimum sum sequence:", findMinSequence(input))
}