package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type UmsAdminRoleRelation struct {
	AdminId   int64      `gorm:"column:admin_id;not null" json:"adminId"`
	RoleId    int64      `gorm:"column:role_id;not null" json:"roleId"`
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	CreateAt  *time.Time `gorm:"column:created_at;not null" json:"createdAt"`
	UpdateAt  *time.Time `gorm:"column:updated_at;not null" json:"updatedAt"`
	DeletedAt *time.Time `gorm:"column:deleted_at;" json:"deletedAt"`

	CreateTime *time.Time            `gorm:"-"`
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"isDel"`
}

func (*UmsAdminRoleRelation) TableName() string {
	return "ums_admin_role_relation"
}
