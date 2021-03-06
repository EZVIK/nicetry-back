package dto

type NiceParams struct {
	Title    string `validate:"required" json:"title"`
	Desc     string `validate:"required" json:"desc"`
	Content  string `validate:"required" json:"content"`
	NiceType uint   `validate:"required" json:"nice_type"`
	Tags     []uint `validate:"required" json:"tags"`
}

// 列表
type INiceParams struct {
	ID       uint   `validate:"required" json:"id"`
	NiceType uint   `validate:"required" json:"nice_type"`
	Title    string `validate:"required" json:"title"`
}

type NiceListParams struct {
	PageSize  int `json:"page_size"`
	PageIndex int `json:"page_index"`
}
