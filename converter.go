package main

import (
	"os"
)

func ConvertFile(f os.FileInfo) FileDTO {
	return FileDTO {f.Name(), f.Size(), !f.IsDir()}
}

func ConvertFiles(files []os.FileInfo) []FileDTO {
	dtos := make([]FileDTO, len(files))
	for i, file := range files {
		dtos[i] = ConvertFile(file)
	} 

	return dtos
}