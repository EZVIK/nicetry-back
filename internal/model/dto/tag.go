package dto

type TagAddParams struct {
	
	Name 		string  `validate:"required" json:"name"`

	ParentId  	uint	`validate:"required" json:"parent_id"`

}
