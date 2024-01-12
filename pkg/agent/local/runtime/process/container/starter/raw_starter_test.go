// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package starter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRawStarter(t *testing.T) {
	c := createTestInitConfig(t)
	s, err := NewRawStarter(c)
	assert.NoError(t, err)
	assert.NoError(t, s.Start())
	assert.NoError(t, s.Wait())
	assert.Equal(t, 0, s.Command().ProcessState.ExitCode())
	assert.NoError(t, s.Release())
}
