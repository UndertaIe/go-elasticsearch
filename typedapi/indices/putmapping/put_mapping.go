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


// Updates the index mappings.
package putmapping

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewPutMapping type alias for index.
type NewPutMapping func(index string) *PutMapping

// NewPutMappingFunc returns a new instance of PutMapping with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutMappingFunc(tp elastictransport.Interface) NewPutMapping {
	return func(index string) *PutMapping {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Updates the index mappings.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-put-mapping.html
func New(tp elastictransport.Interface) *PutMapping {
	r := &PutMapping{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *PutMapping) Raw(raw json.RawMessage) *PutMapping {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutMapping) Request(req *Request) *PutMapping {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutMapping: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_mapping")

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r PutMapping) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutMapping query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the PutMapping headers map.
func (r *PutMapping) Header(key, value string) *PutMapping {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names the mapping should be added to
// (supports wildcards); use `_all` or omit to add the mapping on all indices.
// API Name: index
func (r *PutMapping) Index(v string) *PutMapping {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *PutMapping) AllowNoIndices(b bool) *PutMapping {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *PutMapping) ExpandWildcards(value string) *PutMapping {
	r.values.Set("expand_wildcards", value)

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *PutMapping) IgnoreUnavailable(b bool) *PutMapping {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *PutMapping) MasterTimeout(value string) *PutMapping {
	r.values.Set("master_timeout", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *PutMapping) Timeout(value string) *PutMapping {
	r.values.Set("timeout", value)

	return r
}

// WriteIndexOnly When true, applies mappings only to the write index of an alias or data
// stream
// API name: write_index_only
func (r *PutMapping) WriteIndexOnly(b bool) *PutMapping {
	r.values.Set("write_index_only", strconv.FormatBool(b))

	return r
}
