package main

import (
	. "filemap"
	"fmt"
)

var file_list = make([]*FileMapper, 0)
var FILE_SIZE int = 1024 * 1024

type FileManager struct {
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (f *FileManager) WriteData(data []byte) {
	last_file := f.GetLastFile(false)
	_, err := last_file.PutMessage(data)
	if err == ErrorFileFull {
		last_file = f.GetLastFile(true)
	}
	last_file.PutMessage(data)
}

func (f *FileManager) GetLastFile(isfull bool) *FileMapper {
	fileName := GetNumberFormatInstance().Format(0)

	if len(file_list) <= 0 {
		ff := NewFileMapper(fileName, FILE_SIZE)
		file_list = append(file_list, ff)
		return ff
	} else {
		if isfull {
			fileName = GetNumberFormatInstance().Format((len(file_list)) * FILE_SIZE)
			fmt.Println(fileName)
			ret := NewFileMapper(fileName, FILE_SIZE)
			file_list = append(file_list, ret)
			return ret
		} else {
			return file_list[len(file_list)-1]
		}
	}
}

func main() {
	fm := NewFileManager()
	for i := 0; i < 500000; i++ {
		fm.WriteData([]byte("ni hao!!"))
	}
}
