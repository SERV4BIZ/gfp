package uuid

func NewV4() (string, error) {
	return NewRandom()
}
