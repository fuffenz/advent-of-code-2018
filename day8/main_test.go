package main

import (
	"reflect"
	"testing"
)

func Test_getResults(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 []int
	}{
		{"example",
			args{[]int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2}},
			138,
			66,
			[]int{}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := getResults(tt.args.data)
			if got != tt.want {
				t.Errorf("getResults() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getResults() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("getResults() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
