// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package logstash

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "logstash", asset.ModuleFieldsPri, Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzslU1v4jAQhu/8ilHPsKc95dBLtZUq7Ye0u3dk4onj4sxE/oDm36/sEEhCAGWLeqpP4AnzPpmZl1nBFpsMDCvnhSsXAF57gxk8dFcPCwCJLre69popg8cFABx/AT9YBoMLgEKjkS5L0RWQqHCQNx7f1JiBshzqw81E5mGmUbbj3Qn0ew/0GDsTuijWnuckCYXlCnyJ0CVNb/Cl9+iYrc9XoXNC4SDWsXh886PAFZx4npi80OQST6BVLaxDGYEmhY5Fwh2aSYQtNnu2ch7F3xKTZEoLXCSag/4SfKkd5GwtuppJgudYua+v7eODwvXK1E3MHRHbnMAWciOcg32JFhMr7pA8sNVKk/A4jeRLi2Is+7+Ne6GCbSViGMSGg08cNhBpUgepHmCs7g08w2qd3mOSkDevmM9k3GIDgiTshAkIEjdBqUinT+wnlDMjOsP7kRnnGu48xaezPp314c667Kr5dL/FHmSo6q6XB2kzIdLJ1yYoTev45X5d+ykq7BhagWvaUeiOE9PUI+0MXqgO3i3hWRuP1i3hV/DxJs7UE0vML0yzZ96uNa0rbYx2k4yGSc0D/PaGeUgD5HWFULDtsYImaNUwZ5I3uEgQfxhWErtGdWhnLayopqnmT/Qfb6Ox2iUx6CrkTIVWwSYz3gZaTy6p9+2v1eOQbIAEIf6lb5oe88RCu/bhXwAAAP//akTzxg=="
}
