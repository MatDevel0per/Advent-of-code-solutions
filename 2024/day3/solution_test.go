package main

import (
	"testing"
)

func Test_multiplyInputPart1(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"},
		}, want: 161},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyInputPart1(tt.args.input); got != tt.want {
				t.Errorf("multiplyInputPart1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_multiplyInputPart2(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{[]string{
			"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"},
		}, want: 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplyInputPart2(tt.args.input); got != tt.want {
				t.Errorf("multiplyInputPart2() = %v, want %v", got, tt.want)
			}
		})
	}
}
