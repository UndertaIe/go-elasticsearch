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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// ShardHealthStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/cluster/health/types.ts#L36-L43
type ShardHealthStats struct {
	ActiveShards       int                       `json:"active_shards"`
	InitializingShards int                       `json:"initializing_shards"`
	PrimaryActive      bool                      `json:"primary_active"`
	RelocatingShards   int                       `json:"relocating_shards"`
	Status             healthstatus.HealthStatus `json:"status"`
	UnassignedShards   int                       `json:"unassigned_shards"`
}

// ShardHealthStatsBuilder holds ShardHealthStats struct and provides a builder API.
type ShardHealthStatsBuilder struct {
	v *ShardHealthStats
}

// NewShardHealthStats provides a builder for the ShardHealthStats struct.
func NewShardHealthStatsBuilder() *ShardHealthStatsBuilder {
	r := ShardHealthStatsBuilder{
		&ShardHealthStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardHealthStats struct
func (rb *ShardHealthStatsBuilder) Build() ShardHealthStats {
	return *rb.v
}

func (rb *ShardHealthStatsBuilder) ActiveShards(activeshards int) *ShardHealthStatsBuilder {
	rb.v.ActiveShards = activeshards
	return rb
}

func (rb *ShardHealthStatsBuilder) InitializingShards(initializingshards int) *ShardHealthStatsBuilder {
	rb.v.InitializingShards = initializingshards
	return rb
}

func (rb *ShardHealthStatsBuilder) PrimaryActive(primaryactive bool) *ShardHealthStatsBuilder {
	rb.v.PrimaryActive = primaryactive
	return rb
}

func (rb *ShardHealthStatsBuilder) RelocatingShards(relocatingshards int) *ShardHealthStatsBuilder {
	rb.v.RelocatingShards = relocatingshards
	return rb
}

func (rb *ShardHealthStatsBuilder) Status(status healthstatus.HealthStatus) *ShardHealthStatsBuilder {
	rb.v.Status = status
	return rb
}

func (rb *ShardHealthStatsBuilder) UnassignedShards(unassignedshards int) *ShardHealthStatsBuilder {
	rb.v.UnassignedShards = unassignedshards
	return rb
}
