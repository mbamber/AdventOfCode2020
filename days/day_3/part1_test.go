package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_3"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart1Answer(t *testing.T) {
	expected := 200

	input := inputs.Load(3)
	out, err := main.Part1(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
