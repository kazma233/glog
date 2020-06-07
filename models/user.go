package models

// UserStatus 用户状态
type UserStatus string

var (
	// UserDisable 用户禁用
	UserDisable UserStatus = "DISABLE"
	// UserEnable 用户启用
	UserEnable UserStatus = "ENABLE"
)

type (
	// User 用户实体
	User struct {
		UserID     string     `json:"userId" bson:"userId"`
		Username   string     `json:"username" bson:"username"`
		Password   string     `json:"password" bson:"password"`
		CreateTime LocalTime  `json:"createTime" bson:"createTime"`
		Status     UserStatus `json:"status" bson:"status"`
	}

	// UserLogin 用户登录请求体
	UserLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
