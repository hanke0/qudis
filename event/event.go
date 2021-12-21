// Copyright 2021 ko-han
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
