package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreFormat(t *testing.T) {
	type TestCase struct {
		input  string
		output string
	}

	testCases := []TestCase{
		{
			input:  "{{.ID}}\t{{.Status}}\t{{.Created}}\t{{.Name}}\t{{.Image}}\t{{.State}}",
			output: "{{.ID}}\t{{.Status}}\t{{.Created}}\t{{.Name}}\t{{.Image}}\t{{.State}}\n",
		},
		{
			input:  "{{.ID}}\\t{{.Status}}\\t{{.Created}}\\t{{.Name}}\\t{{.Image}}\\t{{.State}}",
			output: "{{.ID}}\t{{.Status}}\t{{.Created}}\t{{.Name}}\t{{.Image}}\t{{.State}}\n",
		},
		{
			input:  "{{.ID}}\t{{.Status}}\t{{.Created}}\\t{{.Name}}\t{{.Image}}\t{{.State}}\n",
			output: "{{.ID}}\t{{.Status}}\t{{.Created}}\t{{.Name}}\t{{.Image}}\t{{.State}}\n",
		},
	}
	for _, testCase := range testCases {
		result := PreFormat(testCase.input)
		assert.Equal(t, testCase.output, result)
	}
}

func TestLabelToString(t *testing.T) {
	type TestCase struct {
		input  map[string]string
		output string
	}
	testCases := []TestCase{
		{
			input: map[string]string{
				"a": "b",
				"c": "d",
			},
			output: "a=b c=d ",
		},
		{
			input: map[string]string{
				"a": "b",
			},
			output: "a=b ",
		},
		{
			input:  map[string]string{},
			output: "",
		},
	}
	for _, testCase := range testCases {
		result := LabelsToString(testCase.input)
		assert.Equal(t, testCase.output, result)
	}
}
