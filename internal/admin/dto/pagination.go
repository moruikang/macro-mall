// @Author moruikang
// @Date 2025/3/16 17:57:00
// @Desc

package dto

type PageQuery struct {
	PageSize int `json:"pageSize" form:"pageSize" query:"pageSize" binding:"required" default:"1"`
	PageNum  int `json:"pageNum" form:"pageNum"  query:"pageNum" binding:"required" default:"10"`
}

type PageResult struct {
	List interface{} `json:"list"`
	Pagination
}

type Pagination struct {
	PageNum   int   `json:"pageNum"`
	PageSize  int   `json:"pageSize"`
	TotalPage int   `json:"totalPage"`
	Total     int64 `json:"total"`
}

type PublicPageQuery struct {
	PageQuery
	Keyword string `json:"keyword" form:"keyword" query:"keyword"`
}
