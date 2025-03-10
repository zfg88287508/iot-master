package model

// User 用户
type User struct {
	Id       string   `json:"id" xorm:"pk"`
	Username string   `json:"username" xorm:"unique"`
	Name     string   `json:"name,omitempty"`
	Email    string   `json:"email,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	Disabled bool     `json:"disabled,omitempty"`
	Created  Time     `json:"created,omitempty" xorm:"created"`
}

type Me struct {
	User       `xorm:"extends"`
	Privileges []string `json:"privileges"`
}

// Password 密码
type Password struct {
	Id       string `json:"id" xorm:"pk"`
	Password string `json:"password"`
}

type UserHistory struct {
	Id      int64  `json:"id"`
	UserId  string `json:"user_id"`
	Event   string `json:"event"`
	Created Time   `json:"created" xorm:"created"`
}
