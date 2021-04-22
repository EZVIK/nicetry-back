package dto

type LoginParams struct {
	Mail 				string 		`validate:"required" json:"mail"`
	Password     		string 		`validate:"required" json:"password"`
}


type RegisterParams struct {
	Nickname  			string 		`validate:"required" json:"nickname"`
	Mail 				string 		`validate:"required" json:"mail"`
	Password     		string 		`validate:"required" json:"password"`
	ReferralCode		string		`validate:"required" json:"referral_code"`
}

type IUsers []uint

