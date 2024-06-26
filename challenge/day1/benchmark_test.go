package day1

import (
	"testing"

	"github.com/Amr-Shams/aoc2024/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.FromFile())
	}
}
