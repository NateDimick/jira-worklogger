package main

import (
	"testing"
	"time"
)

func TestBuildRequestBody(t *testing.T) {
	type Input struct {
		comment string
		start   *time.Time
		end     *time.Time
	}
	tests := []struct {
		input    Input
		expected AddWorklogBody
	}{
		// todo someday maybe
	}
	for _, tt := range tests {
		t.Run(tt.expected.Comment, func(t *testing.T) {

		})
	}
}

func TestBuildTimeSpent(t *testing.T) {
	tests := []struct {
		seconds  int
		expected string
	}{
		{45, "1m"},
		{150, "2m"},
		{2805, "46m"},
		{3700, "1.0h"},
		{7327, "2.0h"},
		{9000, "2.5h"},
	}
	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := buildTimeSpent(tt.seconds)
			if result != tt.expected {
				t.Log(result)
				t.Fail()
			}
		})
	}
}
