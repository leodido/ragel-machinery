package ragelmachinery

import (
	"bufio"
	"io"
)

// ArbitraryReader returns a Reader that reads from r
// but stops when it finds a delimiter.
// The underlying implementation is a *DelimitedReader.
func ArbitraryReader(r io.Reader) *DelimitedReader {
	return &DelimitedReader{
		reader: bufio.NewReader(r),
		parsingState: parsingState{
			data: []byte{},
			p:    0,
			pe:   0,
			eof:  -1,
		},
	}
}

// DelimitedReader reads arbitrarily sized bytes slices until a delimiter is found.
// It depends on and it keeps track of Ragel's state variables.
type DelimitedReader struct {
	reader *bufio.Reader

	parsingState
}

// Read reads a chunk of bytes until it finds a delimiter.
//
// It always works on the current boundaries of the data,
// and updates them accordingly.
// It returns the number of bytes read and, eventually, an error.
// When delim is not found it returns an io.ErrUnexpectedEOF.
func (r *DelimitedReader) Read(delim byte) (n int, err error) {
	p := r.p

	// Process only the data still to read when P is greater than the half of the data
	// data = a b c d e f g h i l m n
	// vars = - - - - - - - p - pe- -
	if p > len(r.data)/2 {
		copy(r.data, r.data[p:len(r.data)])
		r.p = 0
		r.pe = r.pe - p
		// data = h i l m n f g h i l m n
		// vars = p - pe- - - - - - - - -
		r.data = r.data[0 : len(r.data)-p]
		// data = h i l m n
		// vars = p - pe- -
	}

	// Read until the first occurrence of the delimiter
	line, err := r.reader.ReadBytes(delim)

	// Storing the data up to and including the delimiter
	r.data = append(r.data, line...)
	// Update the position of end
	r.pe = len(r.data)
	if err == io.EOF {
		if len(line) != 0 {
			err = io.ErrUnexpectedEOF
		}
		// Update the position of EOF
		r.eof = r.pe
	}

	return len(line), err
}
