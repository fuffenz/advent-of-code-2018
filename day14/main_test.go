package main

import "testing"

func Test_doStuff(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"example 1", args{9}, "5158916779", 13},
		{"example 2", args{5}, "0124515891", 9},
		{"example 3", args{18}, "9251071085", 48},
		{"example 4", args{2018}, "5941429882", 86764},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := doStuff(tt.args.count)
			if got != tt.want {
				t.Errorf("doStuff() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("doStuff() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
