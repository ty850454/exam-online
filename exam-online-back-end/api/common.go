package api

import (
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

type DateTime time.Time

const timeFormat = "2006-01-02 15:04:05"
const dateFormat = "2006-01-02"

func (t DateTime) MarshalJSON() ([]byte, error) {
	origin := time.Time(t)

	if origin.IsZero() {
		return []byte("null"), nil
	}
	return []byte(origin.Format(timeFormat)), nil
}

// func (t DateTime) UnmarshalJSON(data []byte) (err error) {
// 	if len(data) == 2 {
// 		*t = Date{Time: time.Time{}}
// 		return
// 	}
// 	loc, _ := time.LoadLocation("Asia/Shanghai")
// 	now, err := time.ParseInLocation(`"`+dateFormat+`"`, string(data), loc)
// 	*t = Date{Time: now}
// 	return
// }
