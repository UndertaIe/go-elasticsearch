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

// DelayedDataCheckConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/Datafeed.ts#L119-L130
type DelayedDataCheckConfig struct {
	// CheckWindow The window of time that is searched for late data. This window of time ends
	// with the latest finalized bucket.
	// It defaults to null, which causes an appropriate `check_window` to be
	// calculated when the real-time datafeed runs.
	// In particular, the default `check_window` span calculation is based on the
	// maximum of `2h` or `8 * bucket_span`.
	CheckWindow *Duration `json:"check_window,omitempty"`
	// Enabled Specifies whether the datafeed periodically checks for delayed data.
	Enabled bool `json:"enabled"`
}

// DelayedDataCheckConfigBuilder holds DelayedDataCheckConfig struct and provides a builder API.
type DelayedDataCheckConfigBuilder struct {
	v *DelayedDataCheckConfig
}

// NewDelayedDataCheckConfig provides a builder for the DelayedDataCheckConfig struct.
func NewDelayedDataCheckConfigBuilder() *DelayedDataCheckConfigBuilder {
	r := DelayedDataCheckConfigBuilder{
		&DelayedDataCheckConfig{},
	}

	return &r
}

// Build finalize the chain and returns the DelayedDataCheckConfig struct
func (rb *DelayedDataCheckConfigBuilder) Build() DelayedDataCheckConfig {
	return *rb.v
}

// CheckWindow The window of time that is searched for late data. This window of time ends
// with the latest finalized bucket.
// It defaults to null, which causes an appropriate `check_window` to be
// calculated when the real-time datafeed runs.
// In particular, the default `check_window` span calculation is based on the
// maximum of `2h` or `8 * bucket_span`.

func (rb *DelayedDataCheckConfigBuilder) CheckWindow(checkwindow *DurationBuilder) *DelayedDataCheckConfigBuilder {
	v := checkwindow.Build()
	rb.v.CheckWindow = &v
	return rb
}

// Enabled Specifies whether the datafeed periodically checks for delayed data.

func (rb *DelayedDataCheckConfigBuilder) Enabled(enabled bool) *DelayedDataCheckConfigBuilder {
	rb.v.Enabled = enabled
	return rb
}
