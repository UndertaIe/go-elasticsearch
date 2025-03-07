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

// IndexSettingsLifecycle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L267-L300
type IndexSettingsLifecycle struct {
	// IndexingComplete Indicates whether or not the index has been rolled over. Automatically set to
	// true when ILM completes the rollover action.
	// You can explicitly set it to skip rollover.
	IndexingComplete *bool `json:"indexing_complete,omitempty"`
	// Name The name of the policy to use to manage the index. For information about how
	// Elasticsearch applies policy changes, see Policy updates.
	Name Name `json:"name"`
	// OriginationDate If specified, this is the timestamp used to calculate the index age for its
	// phase transitions. Use this setting
	// if you create a new index that contains old data and want to use the original
	// creation date to calculate the index
	// age. Specified as a Unix epoch value in milliseconds.
	OriginationDate *int64 `json:"origination_date,omitempty"`
	// ParseOriginationDate Set to true to parse the origination date from the index name. This
	// origination date is used to calculate the index age
	// for its phase transitions. The index name must match the pattern
	// ^.*-{date_format}-\\d+, where the date_format is
	// yyyy.MM.dd and the trailing digits are optional. An index that was rolled
	// over would normally match the full format,
	// for example logs-2016.10.31-000002). If the index name doesn’t match the
	// pattern, index creation fails.
	ParseOriginationDate *bool `json:"parse_origination_date,omitempty"`
	// RolloverAlias The index alias to update when the index rolls over. Specify when using a
	// policy that contains a rollover action.
	// When the index rolls over, the alias is updated to reflect that the index is
	// no longer the write index. For more
	// information about rolling indices, see Rollover.
	RolloverAlias *string                     `json:"rollover_alias,omitempty"`
	Step          *IndexSettingsLifecycleStep `json:"step,omitempty"`
}

// IndexSettingsLifecycleBuilder holds IndexSettingsLifecycle struct and provides a builder API.
type IndexSettingsLifecycleBuilder struct {
	v *IndexSettingsLifecycle
}

// NewIndexSettingsLifecycle provides a builder for the IndexSettingsLifecycle struct.
func NewIndexSettingsLifecycleBuilder() *IndexSettingsLifecycleBuilder {
	r := IndexSettingsLifecycleBuilder{
		&IndexSettingsLifecycle{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingsLifecycle struct
func (rb *IndexSettingsLifecycleBuilder) Build() IndexSettingsLifecycle {
	return *rb.v
}

// IndexingComplete Indicates whether or not the index has been rolled over. Automatically set to
// true when ILM completes the rollover action.
// You can explicitly set it to skip rollover.

func (rb *IndexSettingsLifecycleBuilder) IndexingComplete(indexingcomplete bool) *IndexSettingsLifecycleBuilder {
	rb.v.IndexingComplete = &indexingcomplete
	return rb
}

// Name The name of the policy to use to manage the index. For information about how
// Elasticsearch applies policy changes, see Policy updates.

func (rb *IndexSettingsLifecycleBuilder) Name(name Name) *IndexSettingsLifecycleBuilder {
	rb.v.Name = name
	return rb
}

// OriginationDate If specified, this is the timestamp used to calculate the index age for its
// phase transitions. Use this setting
// if you create a new index that contains old data and want to use the original
// creation date to calculate the index
// age. Specified as a Unix epoch value in milliseconds.

func (rb *IndexSettingsLifecycleBuilder) OriginationDate(originationdate int64) *IndexSettingsLifecycleBuilder {
	rb.v.OriginationDate = &originationdate
	return rb
}

// ParseOriginationDate Set to true to parse the origination date from the index name. This
// origination date is used to calculate the index age
// for its phase transitions. The index name must match the pattern
// ^.*-{date_format}-\\d+, where the date_format is
// yyyy.MM.dd and the trailing digits are optional. An index that was rolled
// over would normally match the full format,
// for example logs-2016.10.31-000002). If the index name doesn’t match the
// pattern, index creation fails.

func (rb *IndexSettingsLifecycleBuilder) ParseOriginationDate(parseoriginationdate bool) *IndexSettingsLifecycleBuilder {
	rb.v.ParseOriginationDate = &parseoriginationdate
	return rb
}

// RolloverAlias The index alias to update when the index rolls over. Specify when using a
// policy that contains a rollover action.
// When the index rolls over, the alias is updated to reflect that the index is
// no longer the write index. For more
// information about rolling indices, see Rollover.

func (rb *IndexSettingsLifecycleBuilder) RolloverAlias(rolloveralias string) *IndexSettingsLifecycleBuilder {
	rb.v.RolloverAlias = &rolloveralias
	return rb
}

func (rb *IndexSettingsLifecycleBuilder) Step(step *IndexSettingsLifecycleStepBuilder) *IndexSettingsLifecycleBuilder {
	v := step.Build()
	rb.v.Step = &v
	return rb
}
