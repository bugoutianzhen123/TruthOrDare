package domain

import "time"

// 用户
type User struct {
	Id           uint64         `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	Email        string         `json:"email"`
	Created      time.Time      `json:"created"`
	Updated      time.Time      `json:"updated"`
	PermissionId uint8          `json:"user_permission_id"`
	Permission   UserPermission `gorm:"foreignkey:PermissionId;reference:Id"`
}

type UserPermission struct {
	Id         uint8  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Permission string `json:"permission"` //user 一般用户        super_user 管理员
}

// 聊天记录
type FriendChatHistory struct {
	Id       uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	SenderId uint64    `json:"sender_id"`
	TargetId uint64    `json:"target_id"`
	Text     string    `json:"text"`
	Created  time.Time `json:"created"`
	Deleted  time.Time `json:"deleted"`
	Sender   User      `gorm:"foreignkey:SenderId;references:Id"`
	Target   User      `gorm:"foreignkey:TargetId;references:Id"`
}

// 好友
type Friend struct {
	Id         uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId     uint64    `json:"user_id"`
	FriendId   uint64    `json:"friend_id"`
	FriendName string    `json:"friend_name"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	User       User      `gorm:"foreignkey:UserId;references:Id"`
	Friend     User      `gorm:"foreignkey:FriendId;references:Id"`
}

// 好友申请
type FriendApplication struct {
	Id         uint64        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Message    string        `json:"message"`
	SenderId   uint64        `json:"sender_id"`
	ReceiverId uint64        `json:"receiver_id"`
	StateId    uint64        `json:"state_id"`
	Created    time.Time     `json:"created"`
	Sender     User          `gorm:"foreignkey:SenderId;references:Id"`
	Receiver   User          `gorm:"foreignkey:ReceiverId;references:Id"`
	State      RequestStatus `gorm:"foreignkey:StateId;references:Id"` //申请状态

}

// 申请状态
type RequestStatus struct {
	Id    uint8  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	State string `json:"state"` //申请中  已同意 已接受  已过期
}

// 群聊
type Group struct {
	Id        uint64    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreatorId uint64    `json:"creater_id"`
	Name      string    `json:"name" `
	Created   time.Time `json:"created"`
	Creator   User      `gorm:"foreignkey:CreatorId;references:Id"`
}

// 群申请
type GroupApplication struct {
	Id       uint64        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Message  string        `json:"message"`
	SenderId uint64        `json:"sender_id"`
	GroupId  uint64        `json:"group_id"`
	StateId  uint8         `json:"group_request_status_id"`
	Created  time.Time     `json:"created"`
	Sender   User          `gorm:"foreignkey:SenderId;references:Id"`
	Receiver Group         `gorm:"foreignkey:ReceiverId;references:Id"`
	State    RequestStatus `gorm:"foreignkey:StateId;references:Id"` //申请状态
}

// 群成员
type GroupMember struct {
	Id         uint64        `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	GroupId    uint64        `json:"group_id"`
	UserId     uint64        `json:"user_id"`
	IdentityId uint8         `json:"group_identity_id"`
	Group      Group         `gorm:"foreignkey:GroupId;references:Id"`
	User       User          `gorm:"foreignkey:UserId;references:Id"`
	Created    time.Time     `json:"created"`
	Identity   GroupIdentity `gorm:"foreignkey:IdentityId;references:Id"`
}

// 群权限
type GroupIdentity struct {
	Id       uint8  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Identity string `json:"permission"` //群主 管理员 一般群员
}

// 群历史记录
type GroupChatHistory struct {
	Id       int64     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Created  time.Time `json:"created"`
	SenderId uint64    `json:"sender_id"`
	GroupId  uint64    `json:"group_id"`
	Text     string    `json:"text"`
	Sender   User      `gorm:"foreignkey:SenderId;references:Id"`
	Group    Group     `gorm:"foreignkey:GroupId;references:Id"`
}
