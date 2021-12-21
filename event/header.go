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
