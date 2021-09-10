package models

type Permission struct {
	Role   string `form:"role" json:"role" binding:"required"`
	Path   string `form:"path" json:"path" binding:"required"`
	Method string `form:"method" json:"method" binding:"required"`
}
