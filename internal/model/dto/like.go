package dto


type LikeParams struct {
	PostId       uint   `validate:"required" json:"post_id"`
	LikeType     uint 	`validate:"required" json:"like_type"`
}