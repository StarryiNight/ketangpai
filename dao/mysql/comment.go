package mysql

import (
	"github.com/jmoiron/sqlx"
	"ketangpai/models"

	"go.uber.org/zap"
)

// CreateComment 发表评论
func CreateComment(comment *models.Comment) (err error) {
	sqlStr := `insert into comment(
	comment_id, content, post_id, author_id, parent_id)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, comment.CommentID, comment.Content, comment.PostID,
		comment.AuthorID, comment.ParentID)
	if err != nil {
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

// GetCommentListByIDs 通过id获取帖子回复
func GetCommentListByIDs(ids []string) (commentList []*models.Comment, err error) {
	sqlStr := `select comment_id, content, post_id, author_id, parent_id, create_time
	from comment
	where comment.parent_id in (?)`
	// 动态填充id

	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		zap.L().Error("query comment failed", zap.Error(err))
		return
	}
	// sqlx.In  使用Rebind()重新绑定它
	query = db.Rebind(query)
	err = db.Select(&commentList, query, args...)
	if err != nil {
		zap.L().Error("query comment failed", zap.Error(err))
	}
	return
}
