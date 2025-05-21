package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		str := hello()
		if str != "Hello, World!!!" {
			t.Errorf("Error")
		}
	})
}
