package models

import (
	"time"
)

// Post 发帖属性
type Post struct {
	PostID      int64    `json:"post_id" db:"post_id"`
	Title       string    `json:"title" db:"title"`
	Content     string    `json:"content" db:"content"`
	AuthorId    int64    `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"-" db:"create_time"`
}

// ApiPostDetail 帖子所属信息的结构体
type ApiPostDetail struct {
	*Post
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
}
