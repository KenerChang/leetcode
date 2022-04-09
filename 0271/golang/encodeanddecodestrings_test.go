package encodeanddecodestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	input := []string{"Hello", "World"}

	c := &Codec{}
	assert.ElementsMatch(t, input, c.Decode(c.Encode(input)))

	input = []string{""}
	assert.ElementsMatch(t, input, c.Decode(c.Encode(input)))

	input = []string{"63/Rc", "h", "BmI3FS~J9#vmk", "7uBZ?7*/", "24h+X", "O "}
	assert.ElementsMatch(t, input, c.Decode(c.Encode(input)))

	input = []string{"$", "ab", "%", "$1"}
	assert.ElementsMatch(t, input, c.Decode(c.Encode(input)))
}
