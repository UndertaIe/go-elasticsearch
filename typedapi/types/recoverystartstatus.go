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

// RecoveryStartStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/recovery/types.ts#L91-L96
type RecoveryStartStatus struct {
	CheckIndexTime         *Duration               `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis DurationValueUnitMillis `json:"check_index_time_in_millis"`
	TotalTime              *Duration               `json:"total_time,omitempty"`
	TotalTimeInMillis      DurationValueUnitMillis `json:"total_time_in_millis"`
}

// RecoveryStartStatusBuilder holds RecoveryStartStatus struct and provides a builder API.
type RecoveryStartStatusBuilder struct {
	v *RecoveryStartStatus
}

// NewRecoveryStartStatus provides a builder for the RecoveryStartStatus struct.
func NewRecoveryStartStatusBuilder() *RecoveryStartStatusBuilder {
	r := RecoveryStartStatusBuilder{
		&RecoveryStartStatus{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryStartStatus struct
func (rb *RecoveryStartStatusBuilder) Build() RecoveryStartStatus {
	return *rb.v
}

func (rb *RecoveryStartStatusBuilder) CheckIndexTime(checkindextime *DurationBuilder) *RecoveryStartStatusBuilder {
	v := checkindextime.Build()
	rb.v.CheckIndexTime = &v
	return rb
}

func (rb *RecoveryStartStatusBuilder) CheckIndexTimeInMillis(checkindextimeinmillis *DurationValueUnitMillisBuilder) *RecoveryStartStatusBuilder {
	v := checkindextimeinmillis.Build()
	rb.v.CheckIndexTimeInMillis = v
	return rb
}

func (rb *RecoveryStartStatusBuilder) TotalTime(totaltime *DurationBuilder) *RecoveryStartStatusBuilder {
	v := totaltime.Build()
	rb.v.TotalTime = &v
	return rb
}

func (rb *RecoveryStartStatusBuilder) TotalTimeInMillis(totaltimeinmillis *DurationValueUnitMillisBuilder) *RecoveryStartStatusBuilder {
	v := totaltimeinmillis.Build()
	rb.v.TotalTimeInMillis = v
	return rb
}
