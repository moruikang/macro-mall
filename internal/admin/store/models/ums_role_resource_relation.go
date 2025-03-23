package models

type UmsRoleResourceRelation struct {
	Model
	// 角色ID
	RoleId int64 `gorm:"column:role_id;not null" json:"roleId"`
	// 资源ID
	ResourceId int64 `gorm:"column:resource_id;not null" json:"resourceId"`
}

func (*UmsRoleResourceRelation) TableName() string {
	return "ums_role_resource_relation"
}
