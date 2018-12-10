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
func (s *State) Set(cs, p, pe, eof int) {
	s.currentState, s.p, s.pe, s.eof = cs, p, pe, eof
}

// Get retrieves the state variables of a ragel parser.
func (s *State) Get() (cs, p, pe, eof int, data []byte) {
	return s.currentState, s.p, s.pe, s.eof, s.data
}
