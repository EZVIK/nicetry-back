package dto

type CommentAddParams struct {
	NiceId  uint   `validate:"required" json:"nice_id"`
	Content string `validate:"required" json:"content"`
}

type CommentGetsParams struct {
	NiceId uint `validate:"required" json:"nice_id"`
}
