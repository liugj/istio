// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package framework

// Component is a interface of a test component
type Component interface {
	// GetName return component name
	GetName() string

	// Bringup doing setup for this component
	// Start() is being called in framework.SetUp()
	Start() error

	// Stop stop this component
	// Stop() is being called in framework.TearDown()
	Stop() error

	// IsAlive check if component is alive/running
	IsAlive() (bool, error)

	// Cleanup clean up tmp files and other resource created by this component
	Cleanup() error
}
