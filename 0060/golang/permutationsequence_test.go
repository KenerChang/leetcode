package permutationsequence

import (
	"testing"
)

func TestGetPermutation(t *testing.T) {
	template := "expect %s, but %s"
	n := 3
	k := 3
	target := "213"

	output := getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

	n = 4
	k = 9
	target = "2314"

	output = getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

	n = 4
	k = 1
	target = "1234"

	output = getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

	n = 1
	k = 1
	target = "1"

	output = getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

	n = 9
	k = 62716
	target = "265183794"

	output = getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

	n = 8
	k = 29705
	target = "68245713"

	output = getPermutation(n, k)
	if target != output {
		t.Errorf(template, target, output)
	}

}
