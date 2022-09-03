package textjustification

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		words     []string
		maxWidth  int
		wantLines []string
	}{
		{
			words:     []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth:  16,
			wantLines: []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			words:     []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth:  16,
			wantLines: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			words:     []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth:  20,
			wantLines: []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		{
			words:     []string{"cat", "dog", "fat"},
			maxWidth:  3,
			wantLines: []string{"cat", "dog", "fat"},
		},
		{
			words:     []string{"a", "cat", "dog", "fat"},
			maxWidth:  4,
			wantLines: []string{"a   ", "cat ", "dog ", "fat "},
		},
		{
			words:     []string{"cat", "a", "dog", "fat"},
			maxWidth:  4,
			wantLines: []string{"cat ", "a   ", "dog ", "fat "},
		},
	}
	for _, tt := range tests {
		gotLines := fullJustify(tt.words, tt.maxWidth)

		assert.Equal(t, len(tt.wantLines), len(gotLines))

		for i := range gotLines {
			assert.Equal(t, tt.wantLines[i], gotLines[i])
		}
	}
}
