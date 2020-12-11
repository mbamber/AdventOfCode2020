package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_9"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 4023754

	input := inputs.Load(9)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
