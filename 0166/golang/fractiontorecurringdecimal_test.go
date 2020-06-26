package fractiontorecurringdecimal

import (
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	target := "0.5"
	result := fractionToDecimal(1, 2)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalII(t *testing.T) {
	target := "2"
	result := fractionToDecimal(2, 1)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalIII(t *testing.T) {
	target := "0.(6)"
	result := fractionToDecimal(2, 3)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalIV(t *testing.T) {
	target := "0.625"
	result := fractionToDecimal(5, 8)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalV(t *testing.T) {
	target := "0.(142857)"
	result := fractionToDecimal(1, 7)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalVI(t *testing.T) {
	target := "1.(6)"
	result := fractionToDecimal(5, 3)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalVII(t *testing.T) {
	target := "0.1(6)"
	result := fractionToDecimal(1, 6)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalVIII(t *testing.T) {
	target := "-6.25"
	result := fractionToDecimal(-50, 8)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalIX(t *testing.T) {
	target := "-0.041(6)"
	result := fractionToDecimal(-1, 24)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalX(t *testing.T) {
	target := "-0.041(6)"
	result := fractionToDecimal(1, -24)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalXI(t *testing.T) {
	target := "0.041(6)"
	result := fractionToDecimal(1, 24)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalXII(t *testing.T) {
	target := "0.00(000000465661289042462740251655654056577585848337359161441621040707904997124914069194026549138227660723878669455195477065427143370461252966751355553982241280310754777158628319049732085502639731402098131932683780538602845887105337854867197032523144157689601770377165713821223802198558308923834223016478952081795603341592860749337303449725)"
	result := fractionToDecimal(1, 214748364)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalXIII(t *testing.T) {
	target := "-2147483648"
	result := fractionToDecimal(-2147483648, 1)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalXIV(t *testing.T) {
	target := "-2147483648"
	result := fractionToDecimal(2147483648, -1)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}

func TestFractionToDecimalXV(t *testing.T) {
	target := "0"
	result := fractionToDecimal(0, -1)

	if target != result {
		t.Errorf("expect %s, got %s", target, result)
	}
}
