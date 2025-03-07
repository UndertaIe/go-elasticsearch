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


package types

// SecurityRolesFile type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L303-L307
type SecurityRolesFile struct {
	Dls  bool  `json:"dls"`
	Fls  bool  `json:"fls"`
	Size int64 `json:"size"`
}

// SecurityRolesFileBuilder holds SecurityRolesFile struct and provides a builder API.
type SecurityRolesFileBuilder struct {
	v *SecurityRolesFile
}

// NewSecurityRolesFile provides a builder for the SecurityRolesFile struct.
func NewSecurityRolesFileBuilder() *SecurityRolesFileBuilder {
	r := SecurityRolesFileBuilder{
		&SecurityRolesFile{},
	}

	return &r
}

// Build finalize the chain and returns the SecurityRolesFile struct
func (rb *SecurityRolesFileBuilder) Build() SecurityRolesFile {
	return *rb.v
}

func (rb *SecurityRolesFileBuilder) Dls(dls bool) *SecurityRolesFileBuilder {
	rb.v.Dls = dls
	return rb
}

func (rb *SecurityRolesFileBuilder) Fls(fls bool) *SecurityRolesFileBuilder {
	rb.v.Fls = fls
	return rb
}

func (rb *SecurityRolesFileBuilder) Size(size int64) *SecurityRolesFileBuilder {
	rb.v.Size = size
	return rb
}
