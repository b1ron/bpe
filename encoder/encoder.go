package encoder

// Table is a BPE translation table.
type Table struct {
	m       map[string]string
	encoded string
}
