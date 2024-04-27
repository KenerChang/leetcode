package validnumber

type State interface {
	IsEndState() bool
	Transition(c byte) State // return nil if there's no valid transition
}

var states []State = []State{
	&S0{},
	&S1{},
	&S2{},
	&S3{},
	&S4{},
	&S5{},
	&S6{},
	&S7{},
	&S8{},
}

type S0 struct {
}

func (s *S0) IsEndState() bool {
	return false
}

func (s *S0) Transition(c byte) State {
	if c == '+' || c == '-' {
		return states[1]
	} else if c >= '0' && c <= '9' {
		return states[4]
	} else if c == '.' {
		return states[2]
	}

	return nil
}

type S1 struct{}

func (s *S1) IsEndState() bool {
	return false
}

func (s *S1) Transition(c byte) State {
	if c == '.' {
		return states[2]
	} else if c >= '0' && c <= '9' {
		return states[4]
	}

	return nil
}

type S2 struct{}

func (s *S2) IsEndState() bool {
	return false
}

func (s *S2) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return states[3]
	}

	return nil
}

type S3 struct{}

func (s *S3) IsEndState() bool {
	return true
}

func (s *S3) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return s
	} else if c == 'e' || c == 'E' {
		return states[5]
	}

	return nil
}

type S4 struct{}

func (s *S4) IsEndState() bool {
	return true
}

func (s *S4) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return s
	} else if c == '.' {
		return states[6]
	} else if c == 'e' || c == 'E' {
		return states[5]
	}

	return nil
}

type S5 struct{}

func (s *S5) IsEndState() bool {
	return false
}

func (s *S5) Transition(c byte) State {
	if c == '+' || c == '-' {
		return states[7]
	} else if c >= '0' && c <= '9' {
		return states[8]
	}

	return nil
}

type S6 struct{}

func (s *S6) IsEndState() bool {
	return true
}

func (s *S6) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return states[3]
	} else if c == 'e' || c == 'E' {
		return states[5]
	}

	return nil
}

type S7 struct{}

func (s *S7) IsEndState() bool {
	return false
}

func (s *S7) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return states[8]
	}

	return nil
}

type S8 struct{}

func (s *S8) IsEndState() bool {
	return true
}

func (s *S8) Transition(c byte) State {
	if c >= '0' && c <= '9' {
		return s
	}

	return nil
}

func isNumber(s string) bool {
	// we can use a state machine which has 9 states
	// to sovle this problem.

	var currentState State = states[0]

	for i := range s {
		currentState = currentState.Transition(s[i])

		if currentState == nil {
			return false
		}
	}

	return currentState.IsEndState()
}
