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

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UseSecurity is a security middleware
func UseSecurity(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if SecurityManager().IsAuthenticated(r) {
			h(w, r, ps)
		} else {
			log.Printf("Unauthorized access: %v", r.URL)
			http.Redirect(w, r, "/", 302)
		}
	}
}

// UseCors is a CORS middleware
func UseCors(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h(w, r, ps)
	}
}
