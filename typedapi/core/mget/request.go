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


package mget

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package mget
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/mget/MultiGetRequest.ts#L25-L91
type Request struct {

	// Docs The documents you want to retrieve. Required if no index is specified in the
	// request URI.
	Docs []types.Operation `json:"docs,omitempty"`

	// Ids The IDs of the documents you want to retrieve. Allowed when the index is
	// specified in the request URI.
	Ids *types.Ids `json:"ids,omitempty"`
}

// RequestBuilder is the builder API for the mget.Request
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
		return nil, fmt.Errorf("could not deserialise json into Mget request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Docs(docs []types.OperationBuilder) *RequestBuilder {
	tmp := make([]types.Operation, len(docs))
	for _, value := range docs {
		tmp = append(tmp, value.Build())
	}
	rb.v.Docs = tmp
	return rb
}

func (rb *RequestBuilder) Ids(ids *types.IdsBuilder) *RequestBuilder {
	v := ids.Build()
	rb.v.Ids = &v
	return rb
}
