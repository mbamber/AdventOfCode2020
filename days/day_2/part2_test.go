package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/day_2"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	expected := 530

	input := inputs.Load(2)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
