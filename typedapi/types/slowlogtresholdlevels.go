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

// SlowlogTresholdLevels type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L490-L495
type SlowlogTresholdLevels struct {
	Debug *Duration `json:"debug,omitempty"`
	Info  *Duration `json:"info,omitempty"`
	Trace *Duration `json:"trace,omitempty"`
	Warn  *Duration `json:"warn,omitempty"`
}

// SlowlogTresholdLevelsBuilder holds SlowlogTresholdLevels struct and provides a builder API.
type SlowlogTresholdLevelsBuilder struct {
	v *SlowlogTresholdLevels
}

// NewSlowlogTresholdLevels provides a builder for the SlowlogTresholdLevels struct.
func NewSlowlogTresholdLevelsBuilder() *SlowlogTresholdLevelsBuilder {
	r := SlowlogTresholdLevelsBuilder{
		&SlowlogTresholdLevels{},
	}

	return &r
}

// Build finalize the chain and returns the SlowlogTresholdLevels struct
func (rb *SlowlogTresholdLevelsBuilder) Build() SlowlogTresholdLevels {
	return *rb.v
}

func (rb *SlowlogTresholdLevelsBuilder) Debug(debug *DurationBuilder) *SlowlogTresholdLevelsBuilder {
	v := debug.Build()
	rb.v.Debug = &v
	return rb
}

func (rb *SlowlogTresholdLevelsBuilder) Info(info *DurationBuilder) *SlowlogTresholdLevelsBuilder {
	v := info.Build()
	rb.v.Info = &v
	return rb
}

func (rb *SlowlogTresholdLevelsBuilder) Trace(trace *DurationBuilder) *SlowlogTresholdLevelsBuilder {
	v := trace.Build()
	rb.v.Trace = &v
	return rb
}

func (rb *SlowlogTresholdLevelsBuilder) Warn(warn *DurationBuilder) *SlowlogTresholdLevelsBuilder {
	v := warn.Build()
	rb.v.Warn = &v
	return rb
}
