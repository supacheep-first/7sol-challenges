package services

import (
	"context"
	"piefiredire/proto/summary"
	"strings"
	"unicode"
)

type SummaryServiceServer struct {
	summary.UnimplementedSummaryServiceServer
}

func (s *SummaryServiceServer) GetSummary(ctx context.Context, req *summary.SummaryRequest) (*summary.SummaryResponse, error) {
	beef := make(map[string]int32)
	allowedWords := map[string]bool{
        "t-bone":   true,
        "fatback":  true,
        "pastrami": true,
        "pork":     true,
        "meatloaf": true,
        "jowl":     true,
        "enim":     true,
        "bresaola": true,
    }

	words := strings.FieldsFunc(req.Text, func(r rune) bool {
        return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '-'
    })

    for _, word := range words {
        word = strings.ToLower(word)
        if allowedWords[word] {
            beef[word]++
        }
    }

	return &summary.SummaryResponse{WordCount: beef}, nil
}
