package main

import "testing"

func TestWinningScore(t *testing.T) {
	type args struct {
		numberOfPlayers int
		lastMarble      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"game 1",
			args{10, 1618},
			8317,
		},
		{"game 2",
			args{13, 7999},
			146373,
		},
		{"game 3",
			args{17, 1104},
			2764,
		},
		{"game 4",
			args{21, 6111},
			54718,
		},
		{"game 5",
			args{30, 5807},
			37305,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WinningScore(tt.args.numberOfPlayers, tt.args.lastMarble); got != tt.want {
				t.Errorf("WinningScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
