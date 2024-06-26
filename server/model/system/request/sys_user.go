package request

import model "github.com/flipped-aurora/gin-vue-admin/server/model/system"

// User register structure
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	AuthorityIds []string `json:"authorityIds"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
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
	AuthorityIds []string `json:"authorityIds"`
	ID           uint
}

type ChangeUserInfo struct {
	NickName     string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`
	HeaderImg    string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"`
	AuthorityIds []string             `json:"authorityIds" gorm:"-"`
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
	ID           uint                 `gorm:"primarykey"`
}
