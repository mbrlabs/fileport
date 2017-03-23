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

import "html/template"

var indexTemplate *template.Template
var loginTemplate *template.Template
var errorTemplate *template.Template

func GetIndexTemplate() *template.Template {
	if FileportConfig.RecompileTemplates || indexTemplate == nil {
		indexTemplate, _ = template.ParseFiles("templates/index.html")
	}

	return indexTemplate
}

func GetLoginTemplate() *template.Template {
	if FileportConfig.RecompileTemplates || indexTemplate == nil {
		loginTemplate, _ = template.ParseFiles("templates/login.html")
	}

	return loginTemplate
}

func GetErrorTemplate() *template.Template {
	if FileportConfig.RecompileTemplates || indexTemplate == nil {
		errorTemplate, _ = template.ParseFiles("templates/error.html")
	}

	return errorTemplate
}
