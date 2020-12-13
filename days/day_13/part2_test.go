package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_13"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 535296695251210

	input := inputs.Load(13)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}

func TestEGCD(t *testing.T) {
	cases := map[string]struct {
		a      int
		b      int
		bezot1 int
		bezot2 int
	}{
		"EGCD of 7 and 5": {
			a:      7,
			b:      5,
			bezot1: -2,
			bezot2: 3,
		},
		"EGCD of 5 and 7": {
			a:      5,
			b:      7,
			bezot1: 3,
			bezot2: -2,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			gcd, bezot1, bezot2 := main.EGCD(data.a, data.b)
			assert.Equal(t, 1, gcd)
			assert.Equal(t, data.bezot1, bezot1)
			assert.Equal(t, data.bezot2, bezot2)
		})
	}
}

func TestChineseRemainder(t *testing.T) {
	cases := map[string]struct {
		congruences map[int]int
		expected    int
	}{
		"simple": {
			congruences: map[int]int{
				2: 5,
				3: 7,
			},
			expected: 17,
		},
	}

	for name, data := range cases {
		t.Run(name, func(t *testing.T) {
			out := main.ChineseRemainder(data.congruences)
			assert.Equal(t, data.expected, out)
		})
	}
}
