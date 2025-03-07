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

// GetUserProfileErrors type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/get_user_profile/types.ts#L25-L28
type GetUserProfileErrors struct {
	Count   int64                        `json:"count"`
	Details map[UserProfileId]ErrorCause `json:"details"`
}

// GetUserProfileErrorsBuilder holds GetUserProfileErrors struct and provides a builder API.
type GetUserProfileErrorsBuilder struct {
	v *GetUserProfileErrors
}

// NewGetUserProfileErrors provides a builder for the GetUserProfileErrors struct.
func NewGetUserProfileErrorsBuilder() *GetUserProfileErrorsBuilder {
	r := GetUserProfileErrorsBuilder{
		&GetUserProfileErrors{
			Details: make(map[UserProfileId]ErrorCause, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GetUserProfileErrors struct
func (rb *GetUserProfileErrorsBuilder) Build() GetUserProfileErrors {
	return *rb.v
}

func (rb *GetUserProfileErrorsBuilder) Count(count int64) *GetUserProfileErrorsBuilder {
	rb.v.Count = count
	return rb
}

func (rb *GetUserProfileErrorsBuilder) Details(values map[UserProfileId]*ErrorCauseBuilder) *GetUserProfileErrorsBuilder {
	tmp := make(map[UserProfileId]ErrorCause, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Details = tmp
	return rb
}
