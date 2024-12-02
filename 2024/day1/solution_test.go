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
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3"},
		}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenPart1(tt.args.input); got != tt.want {
				t.Errorf("distanceBetweenPart1() = %v, want %v", got, tt.want)
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
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3"},
		}, want: 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distanceBetweenPart2(tt.args.input); got != tt.want {
				t.Errorf("distanceBetweenPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
