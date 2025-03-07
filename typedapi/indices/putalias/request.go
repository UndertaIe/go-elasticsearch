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


package putalias

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putalias
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/put_alias/IndicesPutAliasRequest.ts#L25-L46
type Request struct {
	Filter *types.QueryContainer `json:"filter,omitempty"`

	IndexRouting *types.Routing `json:"index_routing,omitempty"`

	IsWriteIndex *bool `json:"is_write_index,omitempty"`

	Routing *types.Routing `json:"routing,omitempty"`

	SearchRouting *types.Routing `json:"search_routing,omitempty"`
}

// RequestBuilder is the builder API for the putalias.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putalias request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Filter(filter *types.QueryContainerBuilder) *RequestBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *RequestBuilder) IndexRouting(indexrouting types.Routing) *RequestBuilder {
	rb.v.IndexRouting = &indexrouting
	return rb
}

func (rb *RequestBuilder) IsWriteIndex(iswriteindex bool) *RequestBuilder {
	rb.v.IsWriteIndex = &iswriteindex
	return rb
}

func (rb *RequestBuilder) Routing(routing types.Routing) *RequestBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *RequestBuilder) SearchRouting(searchrouting types.Routing) *RequestBuilder {
	rb.v.SearchRouting = &searchrouting
	return rb
}
