package ragelmachinery

// Reader is the interface that wraps the Read method.
// Its implementations are intended to be used with Ragel parsers or scanners.
type Reader interface {
	Read() (n int, err error)
}
