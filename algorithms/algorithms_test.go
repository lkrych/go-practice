package main

import "testing"

const filepath = "./algorithm_input/mediumUF.txt"

func BenchmarkNaiveUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUF(filepath, &NaiveUF{})
	}
}

func BenchmarkQuickUnionUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUF(filepath, &QuickUnionUF{})
	}
}

func BenchmarkWeightedQuickUnionUF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUF(filepath, &WeightedQuickUnionUF{})
	}
}

func BenchmarkWeightedUnionFindWithCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUF(filepath, &WeightedQuickUnionWithCompressionUF{})
	}
}
