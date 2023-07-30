package api

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
