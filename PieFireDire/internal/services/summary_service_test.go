package services

import (
	"context"
	"piefiredire/proto/summary"
	"reflect"
	"testing"
)

func TestGetSummary(t *testing.T) {
	service := &SummaryServiceServer{}

	tests := []struct {
		name     string
		input    string
		expected map[string]int32
	}{
		{
			name:  "Valid input with allowed words",
			input: "t-bone fatback pastrami pork meatloaf jowl enim bresaola",
			expected: map[string]int32{
				"t-bone":   1,
				"fatback":  1,
				"pastrami": 1,
				"pork":     1,
				"meatloaf": 1,
				"jowl":     1,
				"enim":     1,
				"bresaola": 1,
			},
		},
		{
			name:  "Input with repeated allowed words",
			input: "t-bone t-bone fatback fatback fatback",
			expected: map[string]int32{
				"t-bone":  2,
				"fatback": 3,
			},
		},
		{
			name:     "Input with no allowed words",
			input:    "chicken fish tofu",
			expected: map[string]int32{},
		},
		{
			name:  "Mixed input with allowed and disallowed words",
			input: "t-bone chicken fatback tofu pastrami",
			expected: map[string]int32{
				"t-bone":   1,
				"fatback":  1,
				"pastrami": 1,
			},
		},
		{
			name:  "Input with special characters and mixed casing",
			input: "T-Bone! Fatback, pastrami... Pork? Meatloaf; Jowl: Enim@Bresaola",
			expected: map[string]int32{
				"t-bone":   1,
				"fatback":  1,
				"pastrami": 1,
				"pork":     1,
				"meatloaf": 1,
				"jowl":     1,
				"enim":     1,
				"bresaola": 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &summary.SummaryRequest{Text: tt.input}
			resp, err := service.GetSummary(context.Background(), req)
			if err != nil {
				t.Fatalf("GetSummary returned an error: %v", err)
			}

			if !reflect.DeepEqual(resp.WordCount, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, resp.WordCount)
			}
		})
	}
}
