// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Package aead provides implementations of the AEAD primitive.
package aead

import (
	"github.com/google/tink/go/tink"
)

// Register registers latest AEAD key types and their managers with the Registry.
func Register() error {
	if err := tink.RegisterKeyManager(NewAESGCMKeyManager()); err != nil {
		return err
	}

	if err := tink.RegisterKeyManager(NewChaCha20Poly1305KeyManager()); err != nil {
		return err
	}

	return tink.RegisterKeyManager(NewXChaCha20Poly1305KeyManager())
}