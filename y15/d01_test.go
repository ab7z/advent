package y15

import "testing"

var floorTests = []struct {
	name  string
	input string
	want  int
}{
	{name: "simple 0 a", input: "(())", want: 0},
	{name: "simple 0 b", input: "()()", want: 0},
	{name: "simple 3 a", input: "(((", want: 3},
	{name: "simple 3 b", input: "(()(()(", want: 3},
	{name: "simple 3 c", input: "))(((((", want: 3},
	{name: "simple -1 a", input: "())", want: -1},
	{name: "simple -1 b", input: "))(", want: -1},
	{name: "simple -3 a", input: ")))", want: -3},
	{name: "simple -3 b", input: ")())())", want: -3},
}

func TestCalcFloor(t *testing.T) {
	for _, tt := range floorTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcFloor([]byte(tt.input)); got != tt.want {
				t.Parallel()
				t.Fatalf("%q want: %v got: %v", tt.name, tt.want, got)
			}
		})
	}
}

var floorTests2 = []struct {
	name  string
	input string
	want  int
}{
	{name: "simple a", input: ")", want: 1},
	{name: "simple 0 b", input: "()())", want: 5},
}

func TestCalcFloor2(t *testing.T) {
	for _, tt := range floorTests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcFloor2([]byte(tt.input)); got != tt.want {
				t.Parallel()
				t.Fatalf("%q want: %v got: %v", tt.name, tt.want, got)
			}
		})
	}
}
