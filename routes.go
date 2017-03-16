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
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var INDEX_TEMPLATE template.Template

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl := GetIndexTemplate()
	tmpl.Execute(w, nil)
}

func ListFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read param
	folder := ps.ByName("folder")
	//showHidden := r.URL.Query().Get("hidden") == "true"

	// get files
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	dtos := ConvertFiles(files)

	// convert to json
	data, err := json.Marshal(dtos)
	if err != nil {
		log.Fatal(err)
		return
	}

	// send
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", data)
}
