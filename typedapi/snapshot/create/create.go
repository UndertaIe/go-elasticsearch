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


// Creates a snapshot in a repository.
package create

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
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	repository string
	snapshot   string
}

// NewCreate type alias for index.
type NewCreate func(repository, snapshot string) *Create

// NewCreateFunc returns a new instance of Create with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	return func(repository, snapshot string) *Create {
		n := New(tp)

		n.Repository(repository)

		n.Snapshot(snapshot)

		return n
	}
}

// Creates a snapshot in a repository.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
func New(tp elastictransport.Interface) *Create {
	r := &Create{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Create) Raw(raw json.RawMessage) *Create {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Create) Request(req *Request) *Create {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.repository))
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.snapshot))

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
func (r Create) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Create query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Repository Repository for the snapshot.
// API Name: repository
func (r *Create) Repository(v string) *Create {
	r.paramSet |= repositoryMask
	r.repository = v

	return r
}

// Snapshot Name of the snapshot. Must be unique in the repository.
// API Name: snapshot
func (r *Create) Snapshot(v string) *Create {
	r.paramSet |= snapshotMask
	r.snapshot = v

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Create) MasterTimeout(value string) *Create {
	r.values.Set("master_timeout", value)

	return r
}

// WaitForCompletion If `true`, the request returns a response when the snapshot is complete. If
// `false`, the request returns a response when the snapshot initializes.
// API name: wait_for_completion
func (r *Create) WaitForCompletion(b bool) *Create {
	r.values.Set("wait_for_completion", strconv.FormatBool(b))

	return r
}
