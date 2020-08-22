package tools

import (
	"time"
)

var (
	// SHLoc 上海时区
	SHLoc *time.Location
)

func init() {
	SHLoc = getShangHaiLoc()
}

// getShangHaiLoc 获取上海时区
func getShangHaiLoc() *time.Location {
	cstSh, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}

	return cstSh
}
