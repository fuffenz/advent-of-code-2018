package main

import "testing"

func Test_doStuff(t *testing.T) {
	data, maxWidth := readInput("testinput.txt")
	type args struct {
		data     []byte
		maxWidth int
	}
	tests := []struct {
		name  string
		args  args
		wantx float32
		wanty float32
	}{
		{"example", args{data, maxWidth}, 7, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotx, goty := doStuff(tt.args.data, tt.args.maxWidth); gotx != tt.wantx || goty != tt.wanty {

				t.Errorf("doStuff() = %v, want %v", gotx, tt.wantx)
				t.Errorf("doStuff() = %v, want %v", goty, tt.wanty)

			}
		})
	}
}
