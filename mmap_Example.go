package main

import (
	"fmt"
	// . "midware"
	// "encoding/binary"
	// "time"
	"os"
	"syscall"
	"unsafe"
)

const SHM_NAME = "my_shm"
const SHM_SIZE = 4 * 1024 * 1024

type MyData struct {
	Col1 int
	Col2 int
	Col3 int
}

func main() {
	WriteFile()
}

func WriteFile() {
	//fd, err := C.my_shm_new(C.CString(SHM_NAME))
	f, err := os.Create("TestMMap")
	if err != nil {
		fmt.Println(err)
		return
	}

	//C.ftruncate(fd, SHM_SIZE)
	f.Truncate(SHM_SIZE)

	//ptr, err := C.mmap(nil, SHM_SIZE, \C.PROT_READ|C.PROT_WRITE, C.MAP_SHARED, fd, 0)
	// ptr, err := syscall.Mmap(f.Fd(), 0, SHM_SIZE, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)

	// maxSizeHigh := uint32((off + int64(SHM_SIZE)) >> 32)
	// maxSizeLow := uint32((off + int64(SHM_SIZE)) & 0xFFFFFFFF)

	// h, errno := syscall.CreateFileMapping(syscall.Handle(f.Fd()), nil, syscall.PAGE_READWRITE, maxSizeHigh, maxSizeLow, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	ptr, _ := mmap(f.Fd(), int64(0), SHM_SIZE)
	f.Close()
	var i *int32 = (*int32)(unsafe.Pointer(ptr))
	*i = int32(6)
	var x *int32 = (*int32)(unsafe.Pointer(ptr + uintptr(4)))
	*x = int32(8)
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
func FindShortSuccessor(b []byte) []byte {
	var res []byte
	for _, c := range b {
		if c != 0xff {
			res = append(res, c+1)
			return res
		}
		res = append(res, c)
	}
	return b
}

func FindShortestSeparator(a, b []byte) []byte {
	i, n := 0, len(a)
	if n > len(b) {
		n = len(b)
	}
	for i < n && a[i] == b[i] {
		i++
	}

	if i >= n {
		// Do not shorten if one string is a prefix of the other
	} else if c := a[i]; c < 0xff && c+1 < b[i] {
		r := make([]byte, i+1)
		copy(r, a)
		r[i]++
		return r
	}
	return a
}
