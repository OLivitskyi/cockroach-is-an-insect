package antfarm

import (
	"testing"
)

// Test data structure
type testCase struct {
	name           string
	input          Combination
	expectedOutput string
}

func TestCombinationString(t *testing.T) {
	testCases := []testCase{
		{
			name: "Simple Combination",
			input: Combination{
				paths: [][]*GraphVertex{
					{&GraphVertex{Name: "A"}, &GraphVertex{Name: "B"}},
					{&GraphVertex{Name: "C"}},
				},
			},
			expectedOutput: "{A B}\nC}\n",
		},
		{
			name:           "Empty Combination",
			input:          Combination{},
			expectedOutput: "",
		},
		{
			name: "Single Path",
			input: Combination{
				paths: [][]*GraphVertex{
					{&GraphVertex{Name: "X"}, &GraphVertex{Name: "Y"}, &GraphVertex{Name: "Z"}},
				},
			},
			expectedOutput: "{X Y Z}\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := tc.input.String()
			if output != tc.expectedOutput {
				t.Errorf("Unexpected output.\nExpected: %s\nGot: %s", tc.expectedOutput, output)
			}
		})
	}
}
