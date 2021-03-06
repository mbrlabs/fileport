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
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// static files
	router.ServeFiles(FileportConfig.StaticFilesPrefix+"/*filepath", http.Dir("static"))

	// error pages
	router.NotFound = http.HandlerFunc(NotFoundErrorPage)
	router.MethodNotAllowed = http.HandlerFunc(MethodNotAllowedErrorPage)

	// add routes
	router.GET("/", IndexHandler)
	router.POST("/login", LoginHandler)
	router.GET("/logout", LogoutHandler)
	router.GET("/api/list/*folder", UseSecurity(UseCors(ListFilesHandler)))
	router.GET("/api/get/*path", UseSecurity(UseCors(SendFileHandler)))

	// print info
	url := "http://" + GetLocalIP() + ":" + strconv.Itoa(FileportConfig.Port)
	fmt.Printf("************************************************************\n")
	fmt.Printf("* Fileport\n*\n")
	fmt.Printf("* Url: %v\n", url)
	fmt.Printf("* Password: %v\n", FileportConfig.Password)
	fmt.Printf("************************************************************\n")
	fmt.Println()

	// start server
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(FileportConfig.Port), router))
}
