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

// Fuzziness holds the union for the following types:
//
//	int
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/common.ts#L114-L114
type Fuzziness interface{}

// FuzzinessBuilder holds Fuzziness struct and provides a builder API.
type FuzzinessBuilder struct {
	v Fuzziness
}

// NewFuzziness provides a builder for the Fuzziness struct.
func NewFuzzinessBuilder() *FuzzinessBuilder {
	return &FuzzinessBuilder{}
}

// Build finalize the chain and returns the Fuzziness struct
func (u *FuzzinessBuilder) Build() Fuzziness {
	return u.v
}

func (u *FuzzinessBuilder) Int(int int) *FuzzinessBuilder {
	u.v = &int
	return u
}

func (u *FuzzinessBuilder) String(string string) *FuzzinessBuilder {
	u.v = &string
	return u
}
