package examples

import (
	"fmt"
	"strings"
)

func Example_newlines() {
	multiline := `example 0
example 1
example 2
example X
`

	results := (&newlinesMachine{}).Parse(strings.NewReader(multiline))

	for _, elem := range results {
		fmt.Println("RECV", elem)
	}
	// Output:
	// OnEOF
	// RECV example 0
	// RECV example 1
	// RECV example 2
}
