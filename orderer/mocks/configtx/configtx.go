/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package configtx

import cb "github.com/hyperledger/fabric/protos/common"

// Manager is a mock implementation of configtx.Manager
type Manager struct {
	// ChainIDVal is returned as the result of ChainID()
	ChainIDVal string

	// SequenceVal is returned as the result of Sequence()
	SequenceVal uint64
}

// ConsensusType returns the ConsensusTypeVal
func (cm *Manager) ChainID() string {
	return cm.ChainIDVal
}

// BatchSize returns the BatchSizeVal
func (cm *Manager) Sequence() uint64 {
	return cm.SequenceVal
}

// Apply panics
func (cm *Manager) Apply(configtx *cb.ConfigurationEnvelope) error {
	panic("Unimplemented")
}

// Validate panics
func (cm *Manager) Validate(configtx *cb.ConfigurationEnvelope) error {
	panic("Unimplemented")

}
