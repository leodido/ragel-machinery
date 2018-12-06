package ragelmachinery

import (
	"bufio"
	"io"
)

// ArbitraryReader returns a Reader that reads from r
// but stops when it finds a delimiter.
// The underlying implementation is a *DelimitedReader.
func ArbitraryReader(r io.Reader, delimiter byte) Reader {
	return &DelimitedReader{
		reader:    bufio.NewReader(r),
		delimiter: delimiter,
	}
}

// DelimitedReader reads arbitrarily sized bytes slices until a delimiter is found.
// It depends on and it keeps track of Ragel's state variables.
type DelimitedReader struct {
	reader    *bufio.Reader
	delimiter byte

	parsingState
}

// Read reads an arbitrarily sized chunk of bytes until it finds a delimiter.
func (r *DelimitedReader) Read() (n int, err error) {
	p := r.p
	data := r.data

	// Process only the data still to read when P is greater than the half of the data
	// data = a b c d e f g h i l m n
	// vars = - - - - - - - p - pe- -
	if p > len(data)/2 {
		copy(data, data[p:len(data)])
		r.p = 0
		r.pe = r.pe - p
		// data = h i l m n f g h i l m n
		// vars = p - pe- - - - - - - - -
		data = data[0 : len(data)-p]
		// data = h i l m n
		// vars = p - pe- -
	}

	// Read until the first occurrence of the delimiter
	line, err := r.reader.ReadBytes(r.delimiter)

	// Storing the data up to and including the delimiter
	r.data = append(data, line...)
	// Update the position of end
	r.pe = len(r.data)
	if err == io.EOF {
		if len(line) != 0 {
			err = io.ErrUnexpectedEOF
		}
		// Update the position of EOF
		r.eof = r.pe
	}

	return n, err
}
