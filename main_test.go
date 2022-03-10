package main

import "testing"

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ { // running it a 1000 times
		main()
	}

}
