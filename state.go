package ragelmachinery

type parsingState struct {
	cs, p, pe, eof int    // standard ragel parsing variables
	data           []byte // standard ragel data pointer
}
