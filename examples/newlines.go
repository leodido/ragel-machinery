package examples

import (
	"fmt"
	parser "github.com/leodido/ragel-machinery/parser"
	"io"
)

const prova_start int = 1
const prova_error int = 0

const prova_en_main int = 1

type newlinesMachine struct {
	// define here your support variables for ragel actions
	lines []string
}

// Exec implements the ragel.Parser interface.
func (m *newlinesMachine) Exec(s *parser.State) (int, int) {
	// Tell it to parse from the start for each byte(10) delimited incoming chunk
	cs := 1
	// Retrieve previously stored parsing variables
	_, p, pe, eof, data := s.Get()
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
		case 5:
			goto st_case_5
		case 6:
			goto st_case_6
		case 7:
			goto st_case_7
		case 8:
			goto st_case_8
		case 9:
			goto st_case_9
		case 10:
			goto st_case_10
		case 11:
			goto st_case_11
		}
		goto st_out
	st_case_1:
		if data[p] == 101 {
			goto st2
		}
		goto st0
	st_case_0:
	st0:
		cs = 0
		goto _out
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 120 {
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 97 {
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 109 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 112 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 108 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 101 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 32 {
			goto st9
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
	tr9:

		{
			m.lines = append(m.lines, string(data[:p+1]))
		}

		goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 10 {
			goto st11
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr9
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
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
	_test_eof5:
		cs = 5
		goto _test_eof
	_test_eof6:
		cs = 6
		goto _test_eof
	_test_eof7:
		cs = 7
		goto _test_eof
	_test_eof8:
		cs = 8
		goto _test_eof
	_test_eof9:
		cs = 9
		goto _test_eof
	_test_eof10:
		cs = 10
		goto _test_eof
	_test_eof11:
		cs = 11
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

func (m *newlinesMachine) OnErr() {
	fmt.Println("OnErr")
}

func (m *newlinesMachine) OnEOF() {
	fmt.Println("OnEOF")
}

// Parse composes a new ragel parser for the incoming stream using the current FSM.
func (m *newlinesMachine) Parse(r io.Reader) []string {
	m.lines = []string{}
	p := parser.New(
		parser.ArbitraryReader(r, '\n'), // How to read the stream
		m,                               // How to parse it
		parser.WithStart(1),             // Options
	)
	p.Parse()
	return m.lines
}
