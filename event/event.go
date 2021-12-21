package event

// Event represents a message that may handled in the future.
type Event interface {
	// ID return the event id.
	ID() ID
	// Header gets the header values of event.
	Header() Header
	// Body gets the event body value.
	Body() []byte
	// SetBody replaces event body to a new value.
	SetBody(body []byte)
	// SetID replaces the default event id.
	SetID(id ID)
}

type event struct {
	id     ID
	header Header
	body   []byte
}

var nilBody = []byte{}

func New() Event {
	e := event{
		id:     GetDefaultIDGenerator().New(),
		header: make(Header),
		body:   nilBody,
	}
	return &e
}

func (e *event) ID() ID {
	return e.id
}

func (e *event) Header() Header {
	return e.header
}

func (e *event) Body() []byte {
	return e.body
}

func (e *event) SetBody(body []byte) {
	e.body = body
}

func (e *event) SetID(id ID) {
	e.id = id
}
