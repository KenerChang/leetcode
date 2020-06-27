package excelsheetcolumntitle

import (
	"testing"
)

// func TestConvertToTitle(t *testing.T) {
// 	target := "A"
// 	result := convertToTitle(1)

// 	if target != result {
// 		t.Errorf("expect %s, got %s", target, result)
// 	}
// }

// func TestConvertToTitleII(t *testing.T) {
// 	target := "AB"
// 	result := convertToTitle(28)

// 	if target != result {
// 		t.Errorf("expect %s, got %s", target, result)
// 	}
// }

func TestConvertToTitleIII(t *testing.T) {
	target := "ZY"
	result := convertToTitle(701)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

// func TestConvertToTitleIV(t *testing.T) {
// 	target := ""
// 	result := convertToTitle(0)

// 	if target != result {
// 		t.Errorf("expect %s, got %s", target, result)
// 	}
// }

// func TestConvertToTitleV(t *testing.T) {
// 	target := "AEQTU"
// 	result := convertToTitle(556889)

// 	if target != result {
// 		t.Errorf("expect %s, got %s", target, result)
// 	}
// }

func TestConvertToTitleVI(t *testing.T) {
	target := "AZ"
	result := convertToTitle(52)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

// func TestConvertToTitleVII(t *testing.T) {
// 	target := "BA"
// 	result := convertToTitle(53)

// 	if target != result {
// 		t.Errorf("expect %s, got %s", target, result)
// 	}
// }

func TestConvertToTitleVIII(t *testing.T) {
	target := "Z"
	result := convertToTitle(26)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}
