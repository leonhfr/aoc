package snailfish

import "testing"

func Test_String(t *testing.T) {
	tests := []struct {
		args string
	}{
		{"[1,2]"},
		{"[[1,2],3]"},
		{"[9,[8,7]]"},
		{"[[1,9],[8,5]]"},
		{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]"},
		{"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]"},
		{"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"},
	}

	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := New(tt.args).String(); got != tt.args {
				t.Errorf("String(%v) = %v", tt.args, got)
			}
		})
	}
}

func Test_Magnitude(t *testing.T) {
	tests := []struct {
		args string
		want int
	}{
		{"[9,1]", 29},
		{"[[9,1],[1,9]]", 129},
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
		{"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]", 4140},
	}

	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := New(tt.args).Magnitude(); got != tt.want {
				t.Errorf("String(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_Add(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"1", []string{"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]"}, "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, b := New(tt.args[0]), New(tt.args[1])
			if got := Add(a, b).String(); got != tt.want {
				t.Errorf("String(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
