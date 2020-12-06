package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_5"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 617

	input := inputs.Load(5)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
