# Fileport

Fileport is a web based file browser.   

## Why?
If you are working on more than one machine you find yourself constantly copying files between them
using a usb stick, or even a cloud service (Dropbox, Google Drive, etc.). Fileport allows you to quickly
spawn up a server in your local network, so that you can easily share data between your devices.

You can also use this project as a media server, since it provides everything you need to watch images & videos.

## Build & Run
```
go get github.com/mbrlabs/fileport
cd $GOPATH/src/github.com/mbrlabs/fileport
go run *.go
```

## Features
- Password protection
- Basic read-only file access
- Image gallery
- Video player

Currently works only on Linux (Windows & Mac comming soon).

## TODO
- Well defined API for use in third party apps
- File sorting
- Fuzzy search
- SSH
- Mobile frontend
- File tree