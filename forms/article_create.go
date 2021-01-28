package forms

type ArticleCreate struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Content     string `form:"content" json:"content" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}
