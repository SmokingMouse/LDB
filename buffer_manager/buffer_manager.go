package buffer_manager

import (
	"LDB/consts"
	"LDB/disk_manager"
	"LDB/types"
	"errors"
	"fmt"
)

type BufferManager struct {
	bufSz int
	lru   *Lru
	dm    *disk_manager.DiskManager
	bf    []byte
}

func NewBufferManager(size int, dm *disk_manager.DiskManager) *BufferManager {
	return &BufferManager{
		bufSz: size,
		lru:   NewLru(1024),
		bf:    make([]byte, size*consts.C_Page_Size),
		dm:    dm,
	}
}

func (bm *BufferManager) GetPage(pg types.PageNo) ([]byte, error) {
	if r := bm.lru.Get(pg); r != nil {
		return r.([]byte), nil
	} else if r, ok := bm.dm.Read(pg); ok {
		bm.lru.Set(pg, r)
		return r, nil
	}
	return nil, errors.New(fmt.Sprintf("Fail to get page %d", pg))
}
