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


package samlinvalidate

import (
	"encoding/json"
	"fmt"
)

// Request holds the request body struct for the package samlinvalidate
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/saml_invalidate/Request.ts#L22-L43
type Request struct {

	// Acs The Assertion Consumer Service URL that matches the one of the SAML realm in
	// Elasticsearch that should be used. You must specify either this parameter or
	// the realm parameter.
	Acs *string `json:"acs,omitempty"`

	// QueryString The query part of the URL that the user was redirected to by the SAML IdP to
	// initiate the Single Logout.
	// This query should include a single parameter named SAMLRequest that contains
	// a SAML logout request that is deflated and Base64 encoded.
	// If the SAML IdP has signed the logout request, the URL should include two
	// extra parameters named SigAlg and Signature that contain the algorithm used
	// for the signature and the signature value itself.
	// In order for Elasticsearch to be able to verify the IdP’s signature, the
	// value of the query_string field must be an exact match to the string provided
	// by the browser.
	// The client application must not attempt to parse or process the string in any
	// way.
	QueryString string `json:"query_string"`

	// Realm The name of the SAML realm in Elasticsearch the configuration. You must
	// specify either this parameter or the acs parameter.
	Realm *string `json:"realm,omitempty"`
}

// RequestBuilder is the builder API for the samlinvalidate.Request
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
		return nil, fmt.Errorf("could not deserialise json into Samlinvalidate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Acs(acs string) *RequestBuilder {
	rb.v.Acs = &acs
	return rb
}

func (rb *RequestBuilder) QueryString(querystring string) *RequestBuilder {
	rb.v.QueryString = querystring
	return rb
}

func (rb *RequestBuilder) Realm(realm string) *RequestBuilder {
	rb.v.Realm = &realm
	return rb
}
