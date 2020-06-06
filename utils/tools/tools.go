package tools

import "time"

var (
	// SHLoc 上海时区
	SHLoc *time.Location
)

func init() {
	SHLoc = getShangHaiLoc()
}

// getShangHaiLoc 获取上海时区
func getShangHaiLoc() *time.Location {
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	return cstSh
}
