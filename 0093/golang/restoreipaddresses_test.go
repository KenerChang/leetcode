package restoreipaddresses

import (
	"testing"
)

func verify(s1 []string, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := map[string]bool{}
	m2 := map[string]bool{}

	for idx := range s1 {
		m1[s1[idx]] = true
		m2[s2[idx]] = true
	}

	for _, s := range s1 {
		_, found := m2[s]
		if !found {
			return false
		}
	}

	for _, s := range s2 {
		_, found := m1[s]
		if !found {
			return false
		}
	}

	return true
}

func TestRestoreIpAddressesI(t *testing.T) {
	s := "25525511135"
	target := []string{"255.255.111.35", "255.255.11.135"}

	result := restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "88888888888888"
	target = []string{}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "4234"
	target = []string{"4.2.3.4"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "8111118"
	target = []string{"8.1.11.118", "8.1.111.18", "8.11.1.118", "8.11.11.18", "8.11.111.8", "8.111.1.18", "8.111.11.8", "81.1.1.118", "81.1.11.18", "81.1.111.8", "81.11.1.18", "81.11.11.8", "81.111.1.8"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "42345"
	target = []string{"4.2.3.45", "4.2.34.5", "4.23.4.5", "42.3.4.5"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "111111111111"
	target = []string{"111.111.111.111"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "1111111111111"
	target = []string{}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = ""
	target = []string{}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "010010"
	target = []string{"0.10.0.10", "0.100.1.0"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "10001010"
	target = []string{"100.0.10.10", "100.0.101.0"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}

	s = "0000"
	target = []string{"0.0.0.0"}
	result = restoreIpAddresses(s)
	if !verify(target, result) {
		t.Errorf("expect %+v, got %+v", target, result)
	}
}
