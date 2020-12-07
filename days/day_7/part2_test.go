package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_7"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 5312

	input := inputs.Load(7)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
