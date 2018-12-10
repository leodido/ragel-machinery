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
        m.lines = append(m.lines, string(data[:p+1]))
    }
}

main := ("example"  . ' '? . digit+ @line) . 10;
}%%

%% write data nofinal;

type machine struct{
    // define here your support variables for ragel actions
    lines []string
}

// Exec implements the ragel.Parser interface.
func (m *machine) Exec(s *parser.State) (int, int) {
    // Retrieve previously stored parsing variables
    cs, p, pe, eof, data := s.Get()
    // Inline FSM code here
	%% write exec;
    // Update parsing variables
	s.Set(%%{ write start; }%%, p, pe, eof)
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
