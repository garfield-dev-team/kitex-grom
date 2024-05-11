package model

import (
	"gorm.io/gorm"
)

type Salary struct {
	gorm.Model
	EmployeeId     uint   `json:"employee_id" gorm:"index"`                              // 员工id，foreignKey
	EnterpriseCode string `json:"enterprise_code" column:"enterprise_code" gorm:"index"` // 企业识别码
	Month          string `json:"month" column:"month"`                                  // 月份，例如 2024-04
	Salary         int32  `json:"salary" column:"salary"`                                // 实发工资
}

func (u *Salary) TableName() string {
	return "salary"
}
