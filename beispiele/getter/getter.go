package getter

type MyType struct {
	owner string
}

/*
// GetOwner ist nicht idiomatisch
func (m *MyType) GetOwner() string {
	return m.owner
}
*/

// Owner ist  idiomatisch
func (m *MyType) Owner() string {
	return m.owner
}

// SetOwner ist  idiomatisch
func (m *MyType) SetOwner(o string) {
	m.owner = o
}
