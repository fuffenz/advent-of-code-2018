package main

import "testing"

func Test_getParts(t *testing.T) {
	type args struct {
		claims []Claim
	}
	data := []string{"#1 @ 1,3: 4x4", "#2 @ 3,1: 4x4", "#3 @ 5,5: 2x2"}
	c := parseClaims(data)

	tests := []struct {
		name      string
		args      args
		wantTotal int
		wantID    string
	}{
		{"example", args{c}, 4, "#3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotID := getParts(tt.args.claims)
			if gotTotal != tt.wantTotal {
				t.Errorf("getParts() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotID != tt.wantID {
				t.Errorf("getParts() gotId = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}
