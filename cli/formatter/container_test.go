package formatter

import (
	"testing"

	"github.com/alibaba/pouch/apis/types"

	"github.com/stretchr/testify/assert"
)

func TestNewContainerHeader(t *testing.T) {
	expected := ContainerHeader{
		"Name":    "Name",
		"ID":      "ID",
		"Status":  "Status",
		"Created": "Created",
		"Image":   "Image",
		"Runtime": "Runtime",
		"Command": "Command",
		"ImageID": "ImageID",
		"Labels":  "Labels",
		"State":   "State",
	}
	result := NewContainerHeader()
	assert.Equal(t, result, expected)
}

func TestNewContainerContext(t *testing.T) {
	type TestCase struct {
		container *types.Container
		expected  ContainerContext
	}

	testCases := []TestCase{
		{
			container: &types.Container{
				Command:    "bash",
				Created:    4,
				HostConfig: &types.HostConfig{Runtime: "runc"},
				ID:         "abcdelj8937",
				Image:      "Image123",
				ImageID:    "234567890",
				Labels:     map[string]string{"a": "b", "c": "d"},
				Names:      []string{"nameA", "nameB"},
				State:      "StateA",
				Status:     "StatusB",
			},
			expected: ContainerContext{
				"Name":    "nameA",
				"ID":      "abcdel",
				"Status":  "StatusB",
				"Created": "49 years" + " ago",
				"Image":   "Image123",
				"Runtime": "runc",
				"Command": "bash",
				"ImageID": "234567890",
				"Labels":  "a=b c=d ",
				"State":   "StateA",
			},
		},
		{
			container: &types.Container{
				Command:    "shell",
				Created:    5,
				HostConfig: &types.HostConfig{Runtime: "runv"},
				ID:         "1234567890",
				Image:      "Image456",
				ImageID:    "abcdefg",
				Labels:     map[string]string{},
				Names:      []string{"nameB", "nameA"},
				State:      "StateB",
				Status:     "StatusA",
			},
			expected: ContainerContext{
				"Name":    "nameB",
				"ID":      "123456",
				"Status":  "StatusA",
				"Created": "49 years" + " ago",
				"Image":   "Image456",
				"Runtime": "runv",
				"Command": "shell",
				"ImageID": "abcdefg",
				"Labels":  "",
				"State":   "StateB",
			},
		},
	}
	for _, testCase := range testCases {
		result := NewContainerContext(testCase.container)
		assert.Equal(t, testCase.expected, result)
	}
}
