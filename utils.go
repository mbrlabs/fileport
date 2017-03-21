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
	"net"
	"io/ioutil"
)

type FileType int

const (
	FILE_TYPE_UNKOWN FileType = 0
	FILE_TYPE_IMAGE  FileType = 1
	FILE_TYPE_VIDEO  FileType = 2
	FILE_TYPE_AUDIO  FileType = 3
	FILE_TYPE_TEXT   FileType = 4
	FILE_TYPE_PDF    FileType = 5
	FILE_TYPE_FOLDER FileType = 6
)

func GetFileType(file os.FileInfo) FileType {
	filetype := FILE_TYPE_UNKOWN
	if file.IsDir() {
		filetype = FILE_TYPE_FOLDER
	} else {
		e := strings.ToLower(filepath.Ext(file.Name()))
		if e == ".png" || e == ".jpg" || e == ".jpeg" || e == ".gif" {
			filetype = FILE_TYPE_IMAGE
		} else if e == ".mp4" || e == ".avi" || e == ".webm" || e == ".mkv" || e == ".flv" {
			filetype = FILE_TYPE_VIDEO
		} else if e == ".mp3" || e == ".wav" || e == ".flac" {
			filetype = FILE_TYPE_AUDIO
		} else if e == ".txt" || e == ".go" || e == ".java" || e == ".rs" || e == ".js" || e == ".html" || e == ".css" || e == ".cpp" {
			filetype = FILE_TYPE_TEXT
		} else if e == ".pdf" {
			filetype = FILE_TYPE_PDF
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


func GetFiles(path string, showHidden bool) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return make([]os.FileInfo, 0)
	}

	if(showHidden) {
		return files
	}

	var filtered []os.FileInfo
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			filtered = append(filtered, file)
		}
	}
	return filtered
}