package multiply_strings

import(
	"testing"
)

func TestMultiply(t *testing.T) {
	num1, num2 := "2", "3"
	target := "6"
	result := multiply(num1, num2)

	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}

	num1, num2 = "200", "3"
	target = "600"
	result = multiply(num1, num2)
	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}

	num1, num2 = "200", "300"
	target = "60000"
	result = multiply(num1, num2)
	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}

	num1, num2 = "1111", "9"
	target = "9999"
	result = multiply(num1, num2)
	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}

	num1, num2 = "1111", "0"
	target = "0"
	result = multiply(num1, num2)
	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}

	num1, num2 = "123456789", "987654321"
	target = "121932631112635269"
	result = multiply(num1, num2)
	if target != result {
		t.Errorf("expect: %s, got: %s", target, result)
	}
}

func TestToSlice(t *testing.T) {
	s1 := "5566"
	s1Result := toSlice(s1)
	if !isEqual(s1, s1Result) {
		t.Errorf("transform %s failed, got: %+v", s1, s1Result)
	}

	s2 := "0"
	s2Result := toSlice(s2)
	if !isEqual(s2, s2Result) {
		t.Errorf("transform %s failed, got: %+v", s2, s2Result)
	}

	s3 := "9999990"
	s3Result := toSlice(s3)
	if !isEqual(s3, s3Result) {
		t.Errorf("transform %s failed, got: %+v", s3, s3Result)
	}
}

func isEqual(s string, sli []int) bool {
	leng := len(s)
	if leng != len(sli) {
		return false
	}

	matched := true
	base := 48
	for i := 0; i < leng; i++ {
		targetIdx := leng - 1 - i
		if int(s[i]) - base != sli[targetIdx] {
			matched = false
			break
		}
	}
	return matched
}

func TestToString(t *testing.T) {
	nums1 := []int {8, 7, 2}
	target1 := "278"
	result1 := toString(nums1)

	if target1 != result1 {
		t.Errorf("expect %s, got: %s", target1, result1)
	}
}