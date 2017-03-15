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
	"os"
	"path/filepath"
	"strings"
)

type FileType int
const (
	FILE_TYPE_UNKOWN 	FileType = 0
	FILE_TYPE_IMAGE 	FileType = 1
	FILE_TYPE_VIDEO 	FileType = 2
	FILE_TYPE_TEXT 		FileType = 3
	FILE_TYPE_PDF 		FileType = 4
	FILE_TYPE_FOLDER 	FileType = 5
)

func GetFileType(file os.FileInfo) FileType {
	filetype := FILE_TYPE_UNKOWN
	if file.IsDir() {
		filetype = FILE_TYPE_FOLDER
	} else {
		e := strings.ToLower(filepath.Ext(file.Name()))
		if e == ".png" || e == ".jpg" || e == ".jpeg" || e == ".gif" {
			filetype = FILE_TYPE_IMAGE
		} else if e == ".mp4" || e == ".avi" || e == ".webm" {
			filetype = FILE_TYPE_VIDEO
		} else if e == ".txt" || e == ".go" || e == ".java" || e == ".rs" || e == ".js" || e == ".html" || e == ".css" || e == ".cpp" {
			filetype = FILE_TYPE_TEXT
		} else if e == ".pdf" {
			filetype = FILE_TYPE_PDF
		}
	}

	return filetype
}