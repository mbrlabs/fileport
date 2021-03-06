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
	"crypto/rand"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// FileType used as enum
type FileType int

const (
	// FileTypeUnkown unkown file type
	FileTypeUnkown FileType = 0
	// FileTypeImage any kind of image
	FileTypeImage FileType = 1
	// FileTypeVideo any kind of video
	FileTypeVideo FileType = 2
	// FileTypeAudio any kind of audio file
	FileTypeAudio FileType = 3
	// FileTypeText any text file
	FileTypeText FileType = 4
	// FileTypePDF pdf file
	FileTypePDF FileType = 5
	// FileTypeFolder not a file but a folder
	FileTypeFolder FileType = 6
)

// GetFileType determines file type of file
func GetFileType(file os.FileInfo) FileType {
	filetype := FileTypeUnkown
	if file.IsDir() {
		filetype = FileTypeFolder
	} else {
		e := strings.ToLower(filepath.Ext(file.Name()))
		if e == ".png" || e == ".jpg" || e == ".jpeg" || e == ".gif" {
			filetype = FileTypeImage
		} else if e == ".mp4" || e == ".avi" || e == ".webm" || e == ".mkv" || e == ".flv" || e == ".mov" || e == ".wmv" {
			filetype = FileTypeVideo
		} else if e == ".mp3" || e == ".wav" || e == ".flac" {
			filetype = FileTypeAudio
		} else if e == ".txt" || e == ".go" || e == ".java" || e == ".rs" || e == ".js" || e == ".html" || e == ".css" || e == ".cpp" {
			filetype = FileTypeText
		} else if e == ".pdf" {
			filetype = FileTypePDF
		}
	}

	return filetype
}

// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetFiles returns a list of files in a folder. For every file, the filter function
// determines if the file should be included or not.
func GetFiles(path string, filter func(os.FileInfo) bool) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	result := make([]os.FileInfo, 0)

	if err != nil {
		return result
	}

	for _, file := range files {
		if filter(file) {
			result = append(result, file)
		}
	}

	return result
}

const (
	// AlphaLowerCase alphabet a-z
	AlphaLowerCase = "abcdefghijklmnopqrstuvwxyz"
	// AlphaUpperCase alphabet A-Z
	AlphaUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric alphabet 0-9
	Numeric = "0123456789"
	// Alpha alphabet a-zA-Z
	Alpha = AlphaLowerCase + AlphaUpperCase
	// AlphaNumeric alphabet a-zA-Z0-9
	AlphaNumeric = Alpha + Numeric
)

// RandomString generates a random string
func RandomString(length int, alphabet string) string {
	alphabetLen := byte(len(alphabet))

	// make generate random byte array
	id := make([]byte, length)
	rand.Read(id)

	// replace rand num with char from alphabet
	for i, b := range id {
		id[i] = alphabet[b%alphabetLen]
	}

	return string(id)
}
