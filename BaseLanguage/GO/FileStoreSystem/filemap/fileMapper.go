package filemap

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"syscall"
	"unsafe"
)

var (
	ErrorFileFull = errors.New("the file is full")
)

type FileMapper struct {
	fileName  string
	fileSize  int
	storePos  int
	commitPos int
	vmptr     uintptr
	vmarray   []byte
}

func NewFileMapper(fileName string, fileSize int) *FileMapper {
	fm := &FileMapper{
		fileName:  fileName,
		fileSize:  fileSize,
		storePos:  0,
		commitPos: 0,
		vmptr:     uintptr(0),
	}

	//fd, err := C.my_shm_new(C.CString(SHM_NAME))
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	f.Truncate(int64(fileSize))

	ptr, _ := mmap(f.Fd(), int64(0), fileSize)
	fm.vmptr = ptr
	head := (*reflect.SliceHeader)(unsafe.Pointer(&fm.vmarray))
	head.Data = ptr
	head.Len = fileSize
	head.Cap = fileSize

	f.Close()

	return fm
}

func (f *FileMapper) Map() {

}

func (f *FileMapper) unMap() {

}

func (f *FileMapper) commit() {

}

func (f *FileMapper) PutMessage(str []byte) (int, error) {
	if f.storePos+len(str) > f.fileSize {
		return f.storePos, ErrorFileFull
	}
	copy(f.vmarray[f.storePos:], str)
	f.storePos += len(str)

	// for i := 0; i < len(str); i++ {
	// 	fmt.Println(len(str))
	// 	f.vmarray[f.storePos] = str[i]
	// 	f.storePos += 1
	// }

	return f.storePos, nil
}

func mmap(hfile uintptr, off int64, len int) (uintptr, error) {
	flProtect := uint32(syscall.PAGE_READONLY)
	dwDesiredAccess := uint32(syscall.FILE_MAP_READ)

	flProtect = syscall.PAGE_READWRITE
	dwDesiredAccess = syscall.FILE_MAP_WRITE

	// The maximum size is the area of the file, starting from 0,
	// that we wish to allow to be mappable. It is the sum of
	// the length the user requested, plus the offset where that length
	// is starting from. This does not map the data into memory.
	maxSizeHigh := uint32((off + int64(len)) >> 32)
	maxSizeLow := uint32((off + int64(len)) & 0xFFFFFFFF)
	// TODO: Do we need to set some security attributes? It might help portability.
	h, errno := syscall.CreateFileMapping(syscall.Handle(hfile), nil, flProtect, maxSizeHigh, maxSizeLow, nil)
	if h == 0 {
		return 0, os.NewSyscallError("CreateFileMapping", errno)
	}

	// Actually map a view of the data into memory. The view's size
	// is the length the user requested.
	fileOffsetHigh := uint32(off >> 32)
	fileOffsetLow := uint32(off & 0xFFFFFFFF)
	addr, errno := syscall.MapViewOfFile(h, dwDesiredAccess, fileOffsetHigh, fileOffsetLow, uintptr(len))
	if addr == 0 {
		return 0, os.NewSyscallError("MapViewOfFile", errno)
	}

	return addr, nil
}

// func main() {
// 	f := NewFileMapper("Hello", 1024*1024)
// 	for i := 0; i < 100; i++ {
// 		f.putMessage([]byte("hello world"))
// 	}

// 	fd, err := os.OpenFile("Hello", os.O_RDWR, 0644)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer fd.Close()
// 	fileData, err := ioutil.ReadAll(fd)
// 	fmt.Println(string(fileData))
// }
