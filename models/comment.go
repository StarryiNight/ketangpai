package models

import "time"

// Comment 帖子的属性
type Comment struct {
	PostID     int64    `db:"post_id" json:"post_id"`
	ParentID   int64    `db:"parent_id" json:"parent_id"`
	CommentID  int64    `db:"comment_id" json:"comment_id"`
	AuthorID   int64    `db:"author_id" json:"author_id"`
	Content    string    `db:"content" json:"content"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
}
