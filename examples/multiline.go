package examples

import (
	"fmt"
	parser "github.com/leodido/ragel-machinery/parser"
	"io"
)


const multiline_start int = 1
const multiline_error int = 0

const multiline_en_main int = 1


type multilineMachine struct {
	// define here your support variables for ragel actions
	item  []byte
	items []string
}

// Exec implements the ragel.Parser interface.
func (m *multilineMachine) Exec(s *parser.State) (int, int) {
	// Retrieve previously stored parsing variables
	cs, p, pe, eof, data := s.Get()
	// Inline FSM code here

	{
		if p == pe {
			goto _test_eof
		}
		switch cs {
		case 1:
			goto st_case_1
		case 0:
			goto st_case_0
		case 2:
			goto st_case_2
		case 3:
			goto st_case_3
		case 4:
			goto st_case_4
		}
		goto st_out
	st_case_1:
		if data[p] == 36 {
			goto tr0
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	tr0:

		if len(m.item) > 0 {
			m.items = append(m.items, string(m.item[:len(m.item)-1]))
		}
		// Initialize a new item
		m.item = make([]byte, 0)

		goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 10 {
			goto st3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 10 {
			goto tr3
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	tr3:

		// Collect data each trailer we encounter
		m.item = append(m.item, data...)

		goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch data[p] {
		case 10:
			goto tr3
		case 36:
			goto tr0
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st3
		}
		goto st0
	st_out:
	_test_eof2:
		cs = 2
		goto _test_eof
	_test_eof3:
		cs = 3
		goto _test_eof
	_test_eof4:
		cs = 4
		goto _test_eof

	_test_eof:
		{
		}
	_out:
		{
		}
	}

	// Update parsing variables
	s.Set(cs, p, pe, eof)
	return p, pe
}

func (m *multilineMachine) OnErr() {
	fmt.Println("OnErr")
}

func (m *multilineMachine) OnEOF() {
	fmt.Println("OnEOF")
}

func (m *multilineMachine) OnCompletion() {
	fmt.Println("OnCompletion")
	if len(m.item) > 0 {
		m.items = append(m.items, string(m.item))
	}
}

func (m *multilineMachine) Parse(r io.Reader) {
	m.items = []string{}
	p := parser.New(
		parser.ArbitraryReader(r, 10), // How to read the stream
		m,                             // How to parse it
		parser.WithStart(1),           // Options
	)
	p.Parse()
}
