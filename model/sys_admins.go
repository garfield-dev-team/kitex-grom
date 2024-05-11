package model

import "gorm.io/gorm"

type SysAdmins struct {
	gorm.Model
	Name           string `json:"name" column:"name"`                       // 角色名称
	Password       string `json:"-" column:"password"`                      // 登录密码
	Role           int8   `json:"role" column:"role"`                       // 角色权限
	EnterpriseCode string `json:"enterprise_code" column:"enterprise_code"` // 企业识别码
}

func (u *SysAdmins) TableName() string {
	return "sys_admins"
}
