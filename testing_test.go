package main

import (
	"testing"
)

func TestSimpleSum(t *testing.T) {
	total := Sum(5, 5)

	t.Log("This is a log", total)

	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d", total, 10)
	}
}

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Two positives should come out positive",
			args: args{
				x: 5,
				y: 7,
			},
			want: 12,
		},
		{
			name: "positive and negative",
			args: args{
				x: 5,
				y: -7,
			},
			want: -2,
		},
		{
			name: "test overflow",
			args: args{
				x: 9223372036854775807,
				y: 10,
			},
			want: -9223372036854775799,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
