package models

import (
	"time"
)

type UmsAdmin struct {
	Model
	Username  string     `gorm:"column:username;not null" json:"username"` // 用户名
	Password  string     `gorm:"column:password;not null" json:"password"` // 密码
	Icon      string     `gorm:"column:icon;" json:"icon"`                 // 头像
	Email     string     `gorm:"column:email;" json:"email"`               // 邮箱
	Nickname  string     `gorm:"column:nick_name;" json:"nickName"`        // 昵称
	Note      string     `gorm:"column:note;" json:"note"`                 // 备注信息
	LoginTime *time.Time `gorm:"column:login_time;" json:"loginTime"`      // 最后登录时间
	Status    int64      `gorm:"column:status;default:1" json:"status"`    // 帐号启用状态：0->禁用；1->启用
}

func (*UmsAdmin) TableName() string {
	return "ums_admin"
}
