package examples

import (
    "io"
    "fmt"
    parser "github.com/leodido/ragel-machinery/parser"
)

%%{
machine prova;
action line {
    {
        // Retrieve from the State the data window that matched current line ..
        _, _, _, line := s.Get()
        m.lines = append(m.lines, string(line[:len(line)-1]))
    }
}
main := ("example" . ' '? . digit+) %line . 10;
}%%

%% write data nofinal;

type machine struct{
    // define here your support variables for ragel actions
    lines []string
}

// Exec implements the ragel.Parser interface.
func (m *machine) Exec(s *parser.State) (int, int) {
    // Retrieve the current state
    cs := s.Current()
    // Retrieve previously stored parsing variables
    p, pe, eof, data := s.Get()
    // Inline FSM code here
	%% write exec;
    // Update parsing variables
	s.Set(p, pe, eof)
	return p, pe
}

func (m *machine) OnErr() {
    fmt.Println("OnErr")
}

func (m *machine) OnEOF() {
    fmt.Println("OnEOF")
}

// Parse composes a new ragel parser for the incoming stream using the current FSM.
func Parse(r io.Reader) []string {
    fsm := &machine{}
    fsm.lines = []string{}
    p := parser.New(
        parser.ArbitraryReader(r, '\n'),        // How to read the stream
        fsm,                                    // How to parse it
        parser.WithStart(%%{ write start; }%%), // Options
    )
    p.Parse()
    return fsm.lines
}
