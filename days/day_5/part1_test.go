package main_test

import (
	"testing"

	main "github.com/mbamber/aoc/days/day_5"
	"github.com/stretchr/testify/assert"
)

func TestGetUpper(t *testing.T) {
	cases := map[string]struct {
		char     byte
		expected bool
	}{
		"F is lower": {
			char:     []byte("F")[0],
			expected: false,
		},
		"B is upper": {
			char:     []byte("B")[0],
			expected: true,
		},
		"L is lower": {
			char:     []byte("L")[0],
			expected: false,
		},
		"R is upper": {
			char:     []byte("R")[0],
			expected: true,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out := main.GetUpper(data.char)
			assert.Equal(t, data.expected, out)
		})
	}
}

func TestSplit(t *testing.T) {
	cases := map[string]struct {
		min         int
		max         int
		upper       bool
		expectedMin int
		expectedMax int
	}{
		"lower of 0-127 is 0-63": {
			min:         0,
			max:         127,
			upper:       false,
			expectedMin: 0,
			expectedMax: 63,
		},
		"upper of 0-63 is 32-63": {
			min:         0,
			max:         63,
			upper:       true,
			expectedMin: 32,
			expectedMax: 63,
		},
		"lower of 32-63 is 32-47": {
			min:         32,
			max:         63,
			upper:       false,
			expectedMin: 32,
			expectedMax: 47,
		},
		"upper of 32-47 is 40-47": {
			min:         32,
			max:         47,
			upper:       true,
			expectedMin: 40,
			expectedMax: 47,
		},
		"upper of 40-47 is 44-47": {
			min:         40,
			max:         47,
			upper:       true,
			expectedMin: 44,
			expectedMax: 47,
		},
		"lower of 44-47 is 44-45": {
			min:         44,
			max:         47,
			upper:       false,
			expectedMin: 44,
			expectedMax: 45,
		},
		"uper of 44-45 is 45-45": {
			min:         44,
			max:         45,
			upper:       true,
			expectedMin: 45,
			expectedMax: 45,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			min, max := main.Split(data.min, data.max, data.upper)
			assert.Equal(t, data.expectedMin, min)
			assert.Equal(t, data.expectedMax, max)
		})
	}
}

func TestGetID(t *testing.T) {
	cases := map[string]struct {
		ticket   string
		expected int
	}{
		"BFFFBBFRRR has ID 567": {
			ticket:   "BFFFBBFRRR",
			expected: 567,
		},
		"FFFBBBFRRR has ID 119": {
			ticket:   "FFFBBBFRRR",
			expected: 119,
		},
		"BBFFBBFRLL has ID 820": {
			ticket:   "BBFFBBFRLL",
			expected: 820,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out := main.GetID(data.ticket)
			assert.Equal(t, data.expected, out)
		})
	}
}
