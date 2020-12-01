package main

import (
	"context"
	"flag"

	"log"

	"github.com/mbamber/aoc/inputs"
)

func main() {
	var part int
	flag.IntVar(&part, "p", 1, "part of the day to run (default: 1)")
	flag.Parse()

	input := inputs.Load(1)

	var out interface{}
	var err error

	switch part {
	case 1:
		out, err = Part1(context.Background(), input)
	case 2:
		out, err = Part2(context.Background(), input)
	default:
		log.Fatalf("unknown part of puzzle: %d", part)
	}

	if err != nil {
		log.Fatal(err)
	}
	log.Println(out)
}
