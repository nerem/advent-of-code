package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsGamePossible(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected bool
	}{
		{"Game 1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
		{"Game2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
		{"Game3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
		{"Game4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
		{"Game5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isGamePossible(tt.input, configuration{red: 12, green: 13, blue: 14}))
		})
	}
}

func TestGetLowestPossibleConfiguration(t *testing.T) {
	var tests = []struct {
		name     string
		input    string
		expected configuration
	}{
		{"Game 1", "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", configuration{4, 2, 6}},
		{"Game2", "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", configuration{1, 3, 4}},
		{"Game3", "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", configuration{20, 13, 6}},
		{"Game4", "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", configuration{14, 3, 15}},
		{"Game5", "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", configuration{6, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, getLowestPossibleConfiguration(tt.input))
		})
	}
}

func TestGetPower(t *testing.T) {
	assert.Equal(t, 48, getConfigurationPower(configuration{4, 2, 6}))
}
