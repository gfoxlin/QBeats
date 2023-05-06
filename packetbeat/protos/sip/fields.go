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

package sip

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "sip", asset.ModuleFieldsPri, AssetSip); err != nil {
		panic(err)
	}
}

// AssetSip returns asset data.
// This is the base64 encoded zlib format compressed contents of protos/sip.
func AssetSip() string {
	return "eJzEmcFy4jgQhu88RRf3+AG4TU12q7ixYWavlGK3sSqy5EgyhLffkrCEMBJZKTUhJ3Dxf91u/v5ROU/whqcVKDosADTVDFew3K43ywVAg6qWdNBU8BVs15snNWBNW1oDHpBraCmyRlULmF6tFgAAT8BJjw5p/vRpwBXspRjdlSvymrdC9sS8AfIqRm1qgZakbWldTYqwwqVGLRqcLrkyb3g6Ctn4q1elXlANgisEpYkeldVXM2aPuhNNHvV9RKUn5Zx3LvWVLodOEnXTpwHlUP+iukMJcmpWmJfnOnPyAaWigufAzTc2SKFFLZjTz7GjpJWQdE85YTnsXx2C08Hvl3XlP9aPTNPdtTWuRoQfOrjsqt1c5kL2agUtYQojTau6wz5r1qbl3y9rOCtjgxgVSvOuhGq0FhUDd0LpEqjRxXiDkEU8o5vzSF3jkEX7YRXQIWlQwoGw8eamCWPimMU0AmymdVVzXk0Y29GsAPhJGIP18w1JcI1c7xjyve5mQCb4PvH5TzfbRxX52LVCHols5gETwatxMF8KZt3a1olS4zJm3JE9cv1/l/sbtnaQ9EA0VqWRsznrHxc74Q3kx49rPx1BIb8kisIKyTgKi+TGUlggFk0hOzeiQnYspmqF71Xu0WJrfld5jdEzhSXmHyw8M36yOFDy9R/Ufym5ztVv8bcWVUPVwMhpR3krcrpf/hIwacFol7fo0rU36HDll988krJtN11fNj0xjpIld2C/4Al27m47rtElkLkr7ZBGF0Fqss+labKfg1op+nLb/i1Ff9e4Fl9sXYt/mHl97wX2tZ2nDezRRRb28KSJPT/bxp4dM7LHZlvZY2NmtthcO1tixNDmwElqXe7pn2fAXVu7IsXOdkUeZu7wDgr8bfvXdxwe8otM7iZ01+dhlWyrhxVibg/h2YYP4THPO7iWhKtiulen8IzysrEbYQqKHwOVmPX4x3MnbQr9XgT9Z44jo+5KbP1j1F3C0BYpkbA+n2hlUWBxgFjuw9LD914643RweHT2PntwbJk9NnvXPDa2Z6oZqoJHi8vt88Y9UowhxZGjLMtNQ7bydGpeKihUpoXMR0NBjQkA6+d0kdL5nCt8OiX/SN6Rgwsp6Pl/AXPedDNV0czdJFIDrwXnWBtBlX0kMfyLPnoimZUgTSNR5UX0rMqEiBV6Fc2pLLbsV+BSy2D+cGz9FwAA//9zzI6C"
}
