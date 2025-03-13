package encoder

// optimizations for the encoder package
// 1. scan tokens in the input string and build the frequency map in parallel
// 2. ...

// Table represents a BPE translation table.
type Table struct {
	t       map[string]string
	freq    map[string]int
	encoded string
}

// NewTable returns a new Table.
func NewTable() *Table {
	return &Table{t: make(map[string]string)}
}
