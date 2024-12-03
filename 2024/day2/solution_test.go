package main

import (
	"testing"
)

func Test_distanceBetweenPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9"},
		}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeReactorReadingPart1(tt.args.input); got != tt.want {
				t.Errorf("safeReactorReadingPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_distanceBetweenPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9"},
		}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeReactorReadingPart2(tt.args.input); got != tt.want {
				t.Errorf("safeReactorReadingPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
