package inforce

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

func (cs *InforceService) AddUser(r *http.Request, req *AddUserReq, resp *AddResp) error {
	c := endpoints.NewContext(r)

	u, err := GetUser(c, req.Username)
	if nil == err {
		resp.IsAdded = false
	}

	u = &UserInfo{Username: req.Username, Firstname: req.Firstname, Lastname: req.Lastname}
	if _, err = AddUser(c, u); nil != err {
		resp.IsAdded = false
		return err
	}

	resp.IsAdded = true
	return nil
}

func (cs *InforceService) ValidateUserGetCompanies(r *http.Request, req *LoginReq, resp *LoginResp) error {
	c := endpoints.NewContext(r)

	u, err := GetUser(c, req.Username)
	if nil != err {
		resp.ErrorMsg = "No such user"
		return nil
	}

	companies, err := u.GetCompanies(c)

	if nil != err {
		resp.ErrorMsg = "Could not get list of companies: " + err.Error()
		return err
	}

	resp.Success = true
	resp.Companies = companies

	return nil
}

func (cs *InforceService) GetUserInfo(r *http.Request, req *GetUserInfoReq, resp *GetUserInfoResp) error {
	c := endpoints.NewContext(r)

	u, err := GetUser(c, req.Username)
	if nil != err {
		resp.ErrorMsg = "No such user"
		return err
	}

	resp.Info = u
	return nil
}

func (cs *InforceService) UserCompanyList(r *http.Request, req *UserCompaniesReq, resp *UserCompaniesResp) error {
	c := endpoints.NewContext(r)

	c.Infof("Here")

	u, err := GetUser(c, req.Username)
	if nil != err {
		resp.ErrorMsg = "No such user"
		return err
	}

	companies, err := u.GetCompanies(c)
	if nil != err {
		resp.ErrorMsg = "An error occurred"
		return err
	}

	var uc []string
	for _, v := range companies {
		uc = append(uc, v.Name)
	}
	resp.Companies = uc
	return nil
}
