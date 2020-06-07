package pageable

import "glog/models"

// Result 创建Page响应
func Result(pageNo, pageSize, total int64, data interface{}) *models.Page {
	return &models.Page{
		Total:     total,
		PageTotal: getTotalPage(pageSize, total),
		PageSize:  pageSize,
		PageNo:    pageNo,
		Data:      data,
	}
}

// Empty 空数据
func Empty(pageSize int64) *models.Page {
	return &models.Page{
		Total:     0,
		PageTotal: 1,
		PageSize:  pageSize,
		PageNo:    1,
		Data:      nil,
	}
}

func getTotalPage(size, total int64) int64 {
	if total == 0 {
		return 0
	}

	pageTotal := total / size
	if total%size != 0 {
		return pageTotal + 1
	}

	return pageTotal
}
