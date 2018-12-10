package parser

// State represents the ragel state variables for parsing.
type State parsingState

type parsingState struct {
	currentState int // _start (cs) at time 0
	errorState   int // _error
	finalState   int // _first_final

	p, pe, eof int    // parsing pointers
	data       []byte // data pointer
}

// Set sets the state variables of a ragel parser.
func (s *State) Set(p, pe, eof int) {
	s.p, s.pe, s.eof = p, pe, eof
}

// Get retrieves the state variables of a ragel parser.
func (s *State) Get() (p, pe, eof int, data []byte) {
	return s.p, s.pe, s.eof, s.data
}

// Current returns the current state in which the FSM is.
func (s *State) Current() int {
	return s.currentState
}
