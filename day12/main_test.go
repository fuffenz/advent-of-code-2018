package main

import "testing"

func Test_getResult(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"example",
			args{[]string{
				"initial state: #..#.#..##......###...###",
				"",
				"...## => #",
				"..#.. => #",
				".#... => #",
				".#.#. => #",
				".#.## => #",
				".##.. => #",
				".#### => #",
				"#.#.# => #",
				"#.### => #",
				"##.#. => #",
				"##.## => #",
				"###.. => #",
				"###.# => #",
				"####. => #",
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getResult(tt.args.data)
		})
	}
}
