package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchInputFromURL() ([][]int, error) {
    resp, err := http.Get("https://raw.githubusercontent.com/7-solutions/backend-challenge/main/files/hard.json")
    if err != nil {
        return nil, fmt.Errorf("failed to fetch input: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }

    var input [][]int
    err = json.Unmarshal(body, &input)
    if err != nil {
        return nil, fmt.Errorf("failed to parse JSON: %v", err)
    }

    return input, nil
}


func findBestWay(input [][]int) int {
    bestWay := []int{}


	// copy original input for see best way points.
    original := make([][]int, len(input))
    for i := range input {
        original[i] = make([]int, len(input[i]))
        copy(original[i], input[i])
    }


	// find max value of two values.
    for i := len(input) - 2; i >= 0; i-- {
        for j := 0; j < len(input[i]); j++ {
            input[i][j] += max(input[i+1][j], input[i+1][j+1])
        }
    }
	


	// append best way point path.
    index := 0
    bestWay = append(bestWay, original[0][0])
    for i := 1; i < len(input); i++ {
        if input[i][index] > input[i][index+1] {
            bestWay = append(bestWay, original[i][index])
        } else {
            index++
            bestWay = append(bestWay, original[i][index])
        }
    }

    fmt.Println("Best way:", bestWay)
    return input[0][0]
}

func main(){
    input, err := fetchInputFromURL()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

	sum := findBestWay(input)
	println("Sum of best way:", sum)
}
