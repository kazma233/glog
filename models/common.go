package models

import (
	"encoding/json"
	"glog/utils/tools"
	"time"
)

// IQuery query实体
type IQuery interface {
	InitDate()
}

// Query 查询对象
type Query struct {
	PageNo   int64 `form:"pageNo"`
	PageSize int64 `form:"pageSize"`
}

// InitDate 初始化数据，如果需要的话
func (q *Query) InitDate() {
	if q.PageNo <= 0 {
		q.PageNo = 1
	}

	if q.PageSize <= 0 {
		q.PageSize = 10
	}
}

type (
	// Page 分页对象
	Page struct {
		Total     int64       `json:"total"`     // 总数量
		PageTotal int64       `json:"pageTotal"` // 总页码
		PageSize  int64       `json:"pageSize"`  // 一次查询的数量
		PageNo    int64       `json:"pageNo"`    // 当前页码
		Data      interface{} `json:"list"`      // 数据
	}
)

var _ json.Marshaler = new(LocalTime)
var _ json.Unmarshaler = new(LocalTime)

// LocalTime local time
type LocalTime struct {
	time.Time
}

// MarshalJSON JOSN序列化
func (t LocalTime) MarshalJSON() ([]byte, error) {
	tune := t.Time.In(tools.SHLoc).Format(time.RFC3339)
	return []byte(`"` + tune + `"`), nil
}

// UnmarshalJSON JSON反序列化
func (t *LocalTime) UnmarshalJSON(b []byte) error {
	var t1 time.Time
	err := json.Unmarshal(b, &t1)
	if err != nil {
		return err
	}

	t.Time = t1
	return nil
}
