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


// Returns information about a repository.
package getrepository

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	repositoryMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetRepository struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	repository string
}

// NewGetRepository type alias for index.
type NewGetRepository func() *GetRepository

// NewGetRepositoryFunc returns a new instance of GetRepository with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetRepositoryFunc(tp elastictransport.Interface) NewGetRepository {
	return func() *GetRepository {
		n := New(tp)

		return n
	}
}

// Returns information about a repository.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
func New(tp elastictransport.Interface) *GetRepository {
	r := &GetRepository{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetRepository) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_snapshot")

		method = http.MethodGet
	case r.paramSet == repositoryMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.repository))

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r GetRepository) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetRepository query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetRepository) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the GetRepository headers map.
func (r *GetRepository) Header(key, value string) *GetRepository {
	r.headers.Set(key, value)

	return r
}

// Repository A comma-separated list of repository names
// API Name: repository
func (r *GetRepository) Repository(v string) *GetRepository {
	r.paramSet |= repositoryMask
	r.repository = v

	return r
}

// Local Return local information, do not retrieve the state from master node
// (default: false)
// API name: local
func (r *GetRepository) Local(b bool) *GetRepository {
	r.values.Set("local", strconv.FormatBool(b))

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *GetRepository) MasterTimeout(value string) *GetRepository {
	r.values.Set("master_timeout", value)

	return r
}
