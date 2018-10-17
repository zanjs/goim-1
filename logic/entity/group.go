package entity

// Group 群组
type Group struct {
	Id        int64       `json:"id"`    // 群组id
	Name      string      `json:"name"`  // 组名
	GroupUser []GroupUser `json:"users"` // 群组用户
}

type GroupAdd struct {
	Name    string `json:"name"`     // 群组名称
	UserIds []int  `json:"user_ids"` // 群组成员
}

type GroupUserUpdate struct {
	GroupId int64 `json:"group_id"` // 群组名称
	UserIds []int `json:"user_ids"` // 群组成员
}
