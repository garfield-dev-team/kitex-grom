package model

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EnterpriseCode string `json:"enterprise_code" column:"enterprise_code" gorm:"index"` // 企业识别码
	Name           string `json:"name" column:"name"`                                    // 员工姓名
	Gender         int64  `json:"gender" column:"gender"`                                // 员工性别
	Age            int64  `json:"age" column:"age"`                                      // 员工年龄
	Avatar         string `json:"avatar" column:"avatar"`                                // 员工照片
	Mobile         string `json:"mobile" column:"mobile"`                                // 手机号
	Introduce      string `json:"introduce" column:"introduce"`                          // 备注信息
	IDCardNum      string `json:"id_card_num" column:"id_card_num"`                      // 身份证号
	BankCardNum    string `json:"bank_card_num" column:"bank_card_num"`                  // 银行卡号
}

func (u *Employee) TableName() string {
	return "employee"
}
