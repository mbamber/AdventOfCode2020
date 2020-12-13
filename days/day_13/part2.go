package main

import (
	"context"
	"math"
	"strconv"
	"strings"
)

// Part2 solves the second part of the day's puzzle
func Part2(ctx context.Context, input string) (interface{}, error) {
	lines := strings.Split(input, "\n")
	busses := strings.Split(lines[1], ",")

	offsets := map[int]int{}

	for i, bus := range busses {
		if bus == "x" {
			continue
		}

		busID, err := strconv.Atoi(bus)
		if err != nil {
			return nil, err
		}

		offsets[-i] = busID // Using -i because of modular congruences
	}

	/*
		Remember modular congruence
		(t + i) mod p = 0
		⇒  (t + i) ≡ 0 (mod p)
		⇒  t ≡ -i (mod p) ≡ p - i (mod p)

		i.e.
		t ≡ -offset (mod busID)
	*/
	return ChineseRemainder(offsets), nil
}

// ChineseRemainder uses the chinese remainder theorem of a map from remainders to moduli
//
// 1. Start with a list of remainders ai and a list of moduli pi.
// 2. Compute P = product of all moduli pi.
// 3. For each ai calculate:
//      ni = N/pi
//      yi = modular inverse of ni modulo pi
//      xi = ai*ni*yi
// 4. Finally calculate the solution t = (sum of all xi) mod P.
func ChineseRemainder(congruences map[int]int) int {
	p := 1
	for _, modulus := range congruences {
		p = p * modulus
	}

	x := 0
	for remainder, modulus := range congruences {
		n := p / modulus
		_, inv, _ := EGCD(n, modulus)
		x = (x + remainder*n*inv) % p
	}

	for x < 0 {
		x = x + p
	}
	return x
}

// EGCD computes the extended GCD
func EGCD(a, b int) (gcd, bezot1, bezot2 int) {
	s, oldS := 0, 1
	t, oldT := 1, 0
	r, oldR := b, a

	for r != 0 {
		q := oldR / r
		oldR, r = r, oldR-q*r
		oldS, s = s, oldS-q*s
		oldT, t = t, oldT-q*t
	}

	return int(math.Abs(float64(oldR % b))), oldS, oldT
}
