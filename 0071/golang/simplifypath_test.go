package simplifypath

import (
	"testing"
)

func TestSimplifyPath(t *testing.T) {
	input := "/home/"
	target := "/home"

	output := simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}

	input = "/../"
	target = "/"

	output = simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}

	input = "/home//foo/"
	target = "/home/foo"

	output = simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}

	input = "/a/./b/../../c/"
	target = "/c"

	output = simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}

	input = "/a/../../b/../c//.//"
	target = "/c"

	output = simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}

	input = "/a//b////c/d//././/.."
	target = "/a/b/c"

	output = simplifyPath(input)
	if output != target {
		t.Errorf("expect %s, but got: %s", target, output)
	}
}
