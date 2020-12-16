package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_14"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := int64(3278997609887)

	input := inputs.Load(14)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
