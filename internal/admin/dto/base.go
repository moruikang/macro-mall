// @Author moruikang
// @Date 2025/3/16 17:38:00
// @Desc

package dto

type Base struct {
	Id int64 `json:"id"`
}

type DeleteDTO struct {
	Ids []int64 `json:"ids"`
}
