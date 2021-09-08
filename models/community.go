package models

import "time"

// Community 查询板块简略信息用到的结构体
type Community struct {
	CommunityID   int64 `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

// CommunityDetail 查询板块详细信息用到的结构体
type CommunityDetail struct {
	CommunityID   int64    `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
