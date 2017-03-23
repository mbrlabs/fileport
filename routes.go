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
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if IsAuthenticated(r) {
		tmpl := GetIndexTemplate()
		currentUser, _ := user.Current()
		home := currentUser.HomeDir
		ip := GetLocalIP()
		tmpl.Execute(w, map[string]string{"home": home, "ip": ip})
	} else {
		tmpl := GetLoginTemplate()
		ip := GetLocalIP()
		tmpl.Execute(w, map[string]string{"ip": ip})
	}
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	code := r.PostFormValue("code")

	if code == FileportConfig.Password {
		// generate session id & set cookie
		sessionID := RandomString(64, AlphaNumeric)
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "session", Value: sessionID, Expires: expiration, HttpOnly: false}
		http.SetCookie(w, &cookie)

		// update session store
		SetLogin(sessionID)
		log.Println("Login successful")
	} else {
		log.Println("Login failed")
	}

	http.Redirect(w, r, "/", 302)
}

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sid, err := r.Cookie("session")
	if err == nil && sid != nil {
		SetLogout(sid.Value)
	}
	http.Redirect(w, r, "/", 302)
}

func ListFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read param
	folder := ps.ByName("folder")
	//showHidden := r.URL.Query().Get("hidden") == "true"

	// get files
	files := GetFiles(folder, false)
	if len(files) == 0 {
		http.Error(w, "", 500)
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

func SendFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// open file
	path := ps.ByName("path")
	file, err := os.Open(path)
	defer file.Close()

	// return if file not found
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.ServeFile(w, r, path)
	//http.ServeContent(w, r, file.Name(), time.Now(), file)
}

func GenericError(w http.ResponseWriter, r *http.Request, args map[string]string) {
	tmpl := GetErrorTemplate()
	tmpl.Execute(w, args)
}

func NotFoundErrorPage(w http.ResponseWriter, r *http.Request) {
	GenericError(w, r, map[string]string{"title": "404", "msg": "Not found"})
}

func MethodNotAllowedErrorPage(w http.ResponseWriter, r *http.Request) {
	GenericError(w, r, map[string]string{"title": "405", "msg": "Method not allowed"})
}
