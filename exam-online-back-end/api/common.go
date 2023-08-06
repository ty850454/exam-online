package api

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type PageReq struct {
	PageNum  int `form:"pageNum,default=1"`
	PageSize int `form:"pageSize,default=10"`
}

type PageRes struct {
	Total int64 `json:"total"`
	Data  any   `json:"data"`
}

var EmptyPage = PageRes{
	Total: 0,
	Data:  nil,
}

func (n *PageReq) GetOffset() int {
	return (n.PageNum - 1) * n.PageSize
}

type DateTime struct {
	time.Time
}

const timeFormat = "2006-01-02 15:04:05"
const dateFormat = "2006-01-02"

func (t *DateTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	formatted := fmt.Sprintf("\"%s\"", t.Format(timeFormat))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t *DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan value of time.Time
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
