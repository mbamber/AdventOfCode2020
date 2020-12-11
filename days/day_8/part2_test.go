package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_8"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 1688

	input := inputs.Load(8)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
