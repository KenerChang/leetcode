package addandsearchword

import (
	"testing"
)

func TestStruct(t *testing.T) {
	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")

	if dict.Search("pad") {
		t.Error("expect false, got true")
	}

	if !dict.Search("bad") {
		t.Error("expect true, got false")
	}

	if !dict.Search(".ad") {
		t.Error("expect true, got false")
	}

	if !dict.Search("b..") {
		t.Error("expect true, got false")
	}
}

func TestStructIII(t *testing.T) {
	dict := Constructor()
	dict.AddWord("at")
	dict.AddWord("and")
	dict.AddWord("an")
	dict.AddWord("add")

	if dict.Search("a") {
		t.Error("expect false, got true")
	}

	if dict.Search(".at") {
		t.Error("expect false, got true")
	}

	dict.AddWord("bat")

	if !dict.Search(".at") {
		t.Error("expect true, got false")
	}

	if !dict.Search("an.") {
		t.Error("expect true, got false")
	}

	if dict.Search("a.d.") {
		t.Error("expect false, got true")
	}

	if dict.Search("b.") {
		t.Error("expect false, got true")
	}

	if dict.Search("ab.") {
		t.Error("expect false, got true")
	}

	if dict.Search("a.d.") {
		t.Error("expect false, got true")
	}

	if dict.Search("b.") {
		t.Error("expect false, got true")
	}

	if !dict.Search("a.d") {
		t.Error("expect true, got false")
	}

	if dict.Search(".") {
		t.Error("expect false, got true")
	}
}

func TestStructII(t *testing.T) {
	dict := Constructor()
	dict.AddWord("a")
	dict.AddWord("ab")

	if !dict.Search("a") {
		t.Error("expect true, got false")
	}

	if !dict.Search("a.") {
		t.Error("expect true, got false")
	}

	if !dict.Search("ab") {
		t.Error("expect true, got false")
	}

	if dict.Search(".a") {
		t.Error("expect false, got true")
	}

	if !dict.Search(".b") {
		t.Error("expect true, got false")
	}

	if dict.Search("ab.") {
		t.Error("expect false, got true")
	}

	if !dict.Search(".") {
		t.Error("expect true, got false")
	}

	if !dict.Search("..") {
		t.Error("expect true, got false")
	}
}
