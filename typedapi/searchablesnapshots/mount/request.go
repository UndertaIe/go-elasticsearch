// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package mount

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package mount
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/searchable_snapshots/mount/SearchableSnapshotsMountRequest.ts#L26-L50
type Request struct {
	IgnoreIndexSettings []string `json:"ignore_index_settings,omitempty"`

	Index types.IndexName `json:"index"`

	IndexSettings map[string]interface{} `json:"index_settings,omitempty"`

	RenamedIndex *types.IndexName `json:"renamed_index,omitempty"`
}

// RequestBuilder is the builder API for the mount.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			IndexSettings: make(map[string]interface{}, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Mount request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) IgnoreIndexSettings(ignore_index_settings ...string) *RequestBuilder {
	rb.v.IgnoreIndexSettings = ignore_index_settings
	return rb
}

func (rb *RequestBuilder) Index(index types.IndexName) *RequestBuilder {
	rb.v.Index = index
	return rb
}

func (rb *RequestBuilder) IndexSettings(value map[string]interface{}) *RequestBuilder {
	rb.v.IndexSettings = value
	return rb
}

func (rb *RequestBuilder) RenamedIndex(renamedindex types.IndexName) *RequestBuilder {
	rb.v.RenamedIndex = &renamedindex
	return rb
}
