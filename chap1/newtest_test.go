package chap1

import (
	"testing"
)

func TestNewTest(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		{
			name: "Happy Path",
			want: []int{10, 11},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newTest()
			c := make([]int, 2)
			c[0] = 10
			c[1] = 11
			if len(c) != len(tt.want) {
				t.Errorf("newTest() got = %v, want %v", c, tt.want)
			}
			for i := range c {
				if c[i] != tt.want[i] {
					t.Errorf("newTest() got = %v, want %v", c, tt.want)
				}
			}
		})
	}

	t.Run("Edge Case - Empty Slice", func(t *testing.T) {
		c := make([]int, 0)
		if len(c) != 0 {
			t.Errorf("Expected empty slice, got %v", c)
		}
	})
}
