package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_6"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 3394

	input := inputs.Load(6)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
