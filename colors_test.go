package gocolors

import (
	"fmt"
	"testing"
)

var table = []struct {
	color string
}{
	{color: "black"},
	{color: "red"},
	{color: "green"},
	{color: "yellow"},
	{color: "blue"},
	{color: "magenta"},
	{color: "cyan"},
	{color: "white"},
}

// BenchmarkColor tests the performance of the Color method.
func BenchmarkColor(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("color: %s", v.color), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Color("text rendered with 8 bit colors", v.color, "")
			}
		})
	}
}

// BenchmarkBrightColor tests the performance of the BrightColor method.
func BenchmarkBrightColor(b *testing.B) {
	for _, v := range table {
		b.Run(fmt.Sprintf("bright color: %s", v.color), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BrightColor("text rendered with 16 bit colors", v.color, "")
			}
		})
	}
}
