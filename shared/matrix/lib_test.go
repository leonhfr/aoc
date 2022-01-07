package matrix

import (
	"fmt"
	"testing"
)

func Test_Transpose(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"12\n34", "13\n24"},
		{"123\n456\n789", "147\n258\n369"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			m := IntMatrix(tt.args)
			m.Transpose()
			if got := m.String(); got != tt.want {
				t.Errorf("Transpose(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_Rotate(t *testing.T) {
	tests := []struct {
		args string
		dir  Direction
		want string
	}{
		{"12\n34", Clockwise, "31\n42"},
		{"12\n34", CounterClockwise, "24\n13"},
		{"123\n456\n789", Clockwise, "741\n852\n963"},
		{"123\n456\n789", CounterClockwise, "369\n258\n147"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			m := IntMatrix(tt.args)
			m.Rotate(tt.dir)
			if got := m.String(); got != tt.want {
				t.Errorf("Rotate(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_HorizontalFlip(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"12\n34", "34\n12"},
		{"123\n456\n789", "789\n456\n123"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			m := IntMatrix(tt.args)
			m.HorizontalFlip()
			if got := m.String(); got != tt.want {
				t.Errorf("Rotate(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func Test_VerticalFlip(t *testing.T) {
	tests := []struct {
		args string
		want string
	}{
		{"12\n34", "21\n43"},
		{"123\n456\n789", "321\n654\n987"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			m := IntMatrix(tt.args)
			m.VerticalFlip()
			if got := m.String(); got != tt.want {
				t.Errorf("Rotate(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}
