// Copyright (c) 2017. Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "net/http"
import "sync"

var (
	smInstance *securityManager
	once       sync.Once
)

// SecurityManager handles security
type securityManager struct {
	sessions map[string]bool
}

// SecurityManager returns the securityManager singleton
func SecurityManager() *securityManager {
	if smInstance == nil {
		once.Do(func() {
			smInstance = &securityManager{sessions: make(map[string]bool)}
		})
	}

	return smInstance
}

func (sm *securityManager) IsAuthenticated(r *http.Request) bool {
	if FileportConfig.NoAuth {
		return true
	}

	sid, err := r.Cookie("session")
	if err == nil && sid != nil {
		_, ok := sm.sessions[sid.Value]
		return ok
	}

	return false
}

func (sm *securityManager) Login(sessionID string) {
	sm.sessions[sessionID] = true
}

func (sm *securityManager) Logout(sessionID string) {
	delete(sm.sessions, sessionID)
}
