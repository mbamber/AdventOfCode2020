package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_12"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	expected := 508

	input := inputs.Load(12)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
