package main_test

import (
	"context"
	"testing"

	main "github.com/mbamber/aoc/days/template"
	"github.com/mbamber/aoc/inputs"
	"github.com/stretchr/testify/assert"
)

func TestPart2Answer(t *testing.T) {
	t.Skip("test template, skipping...")
	expected := 0

	input := inputs.Load(-1)
	out, err := main.Part2(context.Background(), input)
	assert.NoError(t, err)
	assert.Equal(t, expected, out)
}
