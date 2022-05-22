package request

import model "github.com/flipped-aurora/gin-vue-admin/server/model/system"

// User register structure
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	Phone        string   `json:"phone"`
	Email        string   `json:"email"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'uploads/file/a81c448a5a9567d71be5e79690c95961_20220425170811.png'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:999"`
	AuthorityIds []string `json:"authorityIds"`
	RefName      string   `json:"refName"`
	RefPhone     string    `json:"refPhone"`
	Region       int       `json:"region"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// User login without structure
type LoginNoCaptcha struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
}


// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId string `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                 `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户角色ID
	AuthorityIds []string             `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}
