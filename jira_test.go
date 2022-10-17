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
		{75, "1m15s"},
		{3700, "1h1m40s"},
		{7327, "2h2m7s"},
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
