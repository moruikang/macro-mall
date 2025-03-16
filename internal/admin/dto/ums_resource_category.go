// @Author moruikang
// @Date 2025/3/16 17:46:00
// @Desc

package dto

type ResourceCategoryCreateDTO struct {
	// 资源分类名称
	Name string `json:"name"`
	// 排序
	Sort int `json:"sort"`
}
