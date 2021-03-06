// Copyright (C) 2017 Google Inc.
//
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

package gles

import (
	"fmt"

	"github.com/google/gapid/core/data/id"
	"github.com/google/gapid/gapis/api"
)

// Name returns the display-name of the context.
func (c Contextʳ) Name() string {
	name := fmt.Sprintf("OpenGL ES context %d", int(c.Identifier()))
	if name := c.Other().ThreadName(); name != "" {
		name += fmt.Sprintf(" - \"%s\"", name)
	}
	return name
}

// ID returns the context's unique identifier.
func (c Contextʳ) ID() api.ContextID {
	if c.IsNil() {
		return api.ContextID{}
	}
	return api.ContextID(id.OfString(fmt.Sprintf("GLES Context %v", c.Identifier())))
}

func (c Contextʳ) Handle() uint32 {
	return uint32(c.Identifier())
}

// API returns the GLES API.
func (c Contextʳ) API() api.API {
	return API{}
}
