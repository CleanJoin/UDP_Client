package main

import "testing"

func TestGetVal(t *testing.T) {
	for i := 0; i < 1000; i++ { // running it a 1000 times
		if start() != 6 {
			t.Error("Incorrect!")
		}
	}

}
