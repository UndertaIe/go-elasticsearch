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

// SyncContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/transform/_types/Transform.ts#L167-L173
type SyncContainer struct {
	// Time Specifies that the transform uses a time field to synchronize the source and
	// destination indices.
	Time *TimeSync `json:"time,omitempty"`
}

// SyncContainerBuilder holds SyncContainer struct and provides a builder API.
type SyncContainerBuilder struct {
	v *SyncContainer
}

// NewSyncContainer provides a builder for the SyncContainer struct.
func NewSyncContainerBuilder() *SyncContainerBuilder {
	r := SyncContainerBuilder{
		&SyncContainer{},
	}

	return &r
}

// Build finalize the chain and returns the SyncContainer struct
func (rb *SyncContainerBuilder) Build() SyncContainer {
	return *rb.v
}

// Time Specifies that the transform uses a time field to synchronize the source and
// destination indices.

func (rb *SyncContainerBuilder) Time(time *TimeSyncBuilder) *SyncContainerBuilder {
	v := time.Build()
	rb.v.Time = &v
	return rb
}
