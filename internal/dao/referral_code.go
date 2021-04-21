package dao

import "nicetry/internal/model"

//func (d *Dao) GetUserReferCodes() (codes model.ReferralCode, err error) {
//
//}

func (d *Dao) GetCode(code string) ( model.ReferralCode, error) {
	rc := model.ReferralCode{Code: code}

	if err := rc.GetCode(d.DB); err != nil {
		return model.ReferralCode{}, err
	}

	return rc, nil
}

func (d *Dao) ConsumeCode(rc *model.ReferralCode) error {
	return nil
}