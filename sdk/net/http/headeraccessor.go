package http

// Interface for accessing http header values.
//
// Go packages use varying data structures and conventions for storing header key-values.
// Using this interface allows us our functions to not be tied to a particular such format.
type HeaderAccessor interface {
	Lookup(key string) []string
}

type headerMapAccessor struct {
	header map[string][]string
}

var _ HeaderAccessor = headerMapAccessor{}

func (accessor headerMapAccessor) Lookup(key string) []string {
	return accessor.header[key]
}
