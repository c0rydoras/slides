package preprocessor

import (
	"reflect"
	"testing"
)

func TestCollectHeadings(t *testing.T) {
	tests := []struct {
		name     string
		slide    []byte
		expected map[int]string
	}{
		{
			name:     "empty slide",
			slide:    []byte(""),
			expected: map[int]string{},
		},
		{
			name:     "no headings",
			slide:    []byte("no headings"),
			expected: map[int]string{},
		},
		{
			name:  "single h1",
			slide: []byte(`# h1`),
			expected: map[int]string{
				1: "h1",
			},
		},
		{
			name: "multiple heading levels",
			slide: []byte(`# h1
## h2
### h3
`),
			expected: map[int]string{
				1: "h1",
				2: "h2",
				3: "h3",
			},
		},
		{
			name: "multiple headings at same level",
			slide: []byte(`
# nono
## h2
# h1`),
			expected: map[int]string{
				1: "h1",
				2: "h2",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := collectHeadings(tt.slide)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("collectHeadings() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAddHeadings(t *testing.T) {
	slides := [][]byte{
		[]byte("# Slide 1\nContent of slide 1"),
		[]byte("## Slide 2\nContent of slide 2"),
		[]byte("Content of slide 3"),
	}

	expected := [][]byte{
		[]byte("# Slide 1\nContent of slide 1"),
		[]byte("# Slide 1\n## Slide 2\nContent of slide 2"),
		[]byte("# Slide 1\n## Slide 2\nContent of slide 3"),
	}

	result := AddHeadings(slides, 2)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("AddHeadings() = %v, want %v", result, expected)
	}
}
