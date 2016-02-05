package main

import (
	"fmt"
	//"time"
	"strconv"
)

/**
简单的分布式ID生成器
**/

type IDGen struct {
	sequence uint32
}

func (ig *IDGen) GetID() string {

}

func main() {
	result := uint64(0)
	// timeLen := uint32(30)
	nodeLen := uint32(10)
	sequenceLen := uint32(24)
	//time_mask := 0x3fffffff
	//node_mask := 0x3ff
	//sequence_mask := 0xffffff

	var time uint64 = 1004653969
	var node uint64 = 123
	var sequence uint64 = 12345

	result = ((time & 0x3fffffff) << (nodeLen + sequenceLen)) | ((node & 0x3ff) << sequenceLen) | (sequence & 0xffffff)

}
