// @Author moruikang
// @Date 2025/3/16 17:46:00
// @Desc

package dto

type ResourceCreateDTO struct {
	// 资源名称
	Name string `json:"name"`
	// 资源URL
	Url string `json:"url"`
	// 描述
	Description string `json:"description"`
	// 资源分类ID
	CategoryId int64 `json:"categoryId"`
}

type ResourcePageQueryDTO struct {
	PageQuery
	// 资源分类ID
	CategoryId int64 `form:"categoryId"`
	// 资源名称模糊关键字
	NameKeyword string `form:"nameKeyword"`
	// 资源URL模糊关键字
	UrlKeyword string `form:"urlKeyword"`
}
