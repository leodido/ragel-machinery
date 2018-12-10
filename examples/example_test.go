package examples

import (
	"fmt"
	"strings"
)

func ExampleParse() {
	multiline := `example 0
example 1
example 2
`

	for _, elem := range Parse(strings.NewReader(multiline)) {
		fmt.Println("RECV", elem)
	}
	// Output:
	// OnEOF
	// RECV example 0
	// RECV example 1
	// RECV example 2
}
