package main

import (
	"testing"
)

func Test_getCheckSum(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example",
			args{[]string{"abcdef", "bababc", "abbcde", "abcccd", "aabcdd", "abcdee", "ababab"}},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCheckSum(tt.args.data); got != tt.want {
				t.Errorf("getCheckSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSmallestDiff(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"example",
			args{[]string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"}},
			"fgij",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSmallestDiff(tt.args.data); got != tt.want {
				t.Errorf("getSmallestDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
