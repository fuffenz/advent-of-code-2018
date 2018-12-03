package main

import (
	"testing"
)

func Test_getResult(t *testing.T) {
	type args struct {
		data []int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
	}{
		{
			"example 1",
			args{[]int64{1, -2, 3, 1}},
			3,
		},
		{
			"example 2",
			args{[]int64{1, 1, 1}},
			3,
		},
		{
			"example 3",
			args{[]int64{1, 1, -2}},
			0,
		},
		{
			"example 4",
			args{[]int64{-1, -2 - 3}},
			-6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getResult(tt.args.data); gotResult != tt.wantResult {
				t.Errorf("getResult() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_getFirstDuplicate(t *testing.T) {
	type args struct {
		data []int64
	}
	tests := []struct {
		name          string
		args          args
		wantDuplicate int64
	}{
		{
			"example 1",
			args{[]int64{3, 3, 4, -2, -4}},
			10,
		},
		{
			"example 2",
			args{[]int64{-6, 3, 8, 5, -6}},
			5,
		},
		{
			"example 3",
			args{[]int64{7, 7, -2, -7, -4}},
			14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDuplicate := getFirstDuplicate(tt.args.data); gotDuplicate != tt.wantDuplicate {
				t.Errorf("getFirstDuplicate() = %v, want %v", gotDuplicate, tt.wantDuplicate)
			}
		})
	}
}
