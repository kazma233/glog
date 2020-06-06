package models

import (
	"encoding/json"
	"glog/utils/tools"
	"time"
)

// Query 查询对象
type Query struct {
	PageNo   int64 `form:"pageNo"`
	PageSize int64 `form:"pageSize"`
}

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
