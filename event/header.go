package event

// Header represents a map of key-values of an event.
type Header map[string]string

// Get gets the string value to the mapping key.
// If there are no associated with the given key,
// it returns an empty string.
func (h Header) Get(key string) string {
	return h[key]
}

// Set sets the string value to the mapping key.
// It replaces any existing value.
func (h Header) Set(key, value string) {
	h[key] = value
}

// Has checks whether a key is set.
func (h Header) Has(key string) bool {
	_, ok := h[key]
	return ok
}

// Del deletes the value associated with the given key.
func (h Header) Del(key string) {
	delete(h, key)
}
