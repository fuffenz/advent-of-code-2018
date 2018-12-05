package main

import "testing"

func Test_getResults(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"example",
			args{"dabAcCaCBAcCcaDA"},
			10,
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getResults(tt.args.s)
			if got != tt.want {
				t.Errorf("getResults() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getResults() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
