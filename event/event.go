package event

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	mrand "math/rand"
	"sync"
)

// ID is a unique identity of a event.
type ID [16]byte

type TraceID [16]byte

var nilID ID
var _ json.Marshaler = nilID

// IsValid checks whether the ID is valid. A valid ID does
// not consist of zeros only.
func (t ID) IsValid() bool {
	return !bytes.Equal(t[:], nilID[:])
}

// MarshalJSON implements a custom marshal function to encode ID
// as a hex string.
func (t ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// String returns the hex string representation form of a ID
func (t ID) String() string {
	return hex.EncodeToString(t[:])
}

// IDGenerator provides an interface to generate event id.
type IDGenerator interface {
	// New returns a new evnet id.
	New() ID
}

type randomIDGenerator struct {
	mu sync.Mutex
	r  *mrand.Rand
}

func newRandomIDGenerator() IDGenerator {
	gen := &randomIDGenerator{}
	var seed int64
	_ = binary.Read(rand.Reader, binary.LittleEndian, &seed)
	gen.r = mrand.New(mrand.NewSource(seed))
	return gen
}

// New returns a non-zero ID from a randomly-chosen sequence.
func (r *randomIDGenerator) New() (id ID) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := rand.Read(id[:])
	if err != nil {
		r.r.Read(id[:])
	}
	return
}
