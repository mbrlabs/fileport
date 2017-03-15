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
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

func main() {
	router := httprouter.New()

	// static files
	router.ServeFiles(FileportConfig.StaticFilesPrefix+"/*filepath", http.Dir("static"))

	// add routes
	router.GET("/", Index)
	router.GET("/api/list/*folder", ListFiles)

	// start server
	log.Printf("Listening on http://localhost:%v", FileportConfig.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(FileportConfig.Port), router))
}
