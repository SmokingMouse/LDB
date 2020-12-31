package disk_manager

import (
	"LDB/consts"
	"LDB/types"
	"log"
	"os"
)

type DiskManager struct {
	handler *os.File
	dbFile  string
}

func NewDBFile(filename string) *DiskManager {
	handler, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Errorf("unable to create db file,%s. %v", filename, err)
	}
	return &DiskManager{
		handler,
		filename,
	}
}

func offset(i types.PageNo) int64 {
	return int64(int(i) * consts.C_Page_Size)

}

func (dm *DiskManager) Read(idx types.PageNo) ([]byte, bool) {
	var data []byte
	succ := true
	_, err := dm.handler.ReadAt(data, offset(idx))
	if err != nil {
		log.Errorf("fail to read data,%v", err)
		succ = false
	}
	return data, succ
}

func (dm *DiskManager) Write(idx types.PageNo, data []byte) bool {
	_, err := dm.handler.WriteAt(data, offset(idx))
	succ := true
	if err != nil {
		log.Errorf("fail to write data,%v", err)
		succ = false
	}
	return succ
}

func (dm *DiskManager) SyncWrite(idx types.PageNo, data []byte) {
	if succ := dm.Write(idx, data); succ {
		dm.handler.Sync()
	}
}
