package inforce

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

type InforceService struct {
}

func (cs *InforceService) CompanyList(r *http.Request, req *CompaniesListReq, resp *CompaniesList) error {
	c := endpoints.NewContext(r)
	companies, err := GetAllCompanies(c)

	if nil != err {
		return err
	}

	resp.Companies = companies
	return nil
}

func (cs *InforceService) AddCompany(
	r *http.Request, req *AddCompanyReq, resp *AddResp) error {
	c := endpoints.NewContext(r)

	company := &Company{Name: req.Name}
	_, err := AddCompany(c, company)

	if nil != err {
		resp.IsAdded = false
		return err
	}

	resp.IsAdded = true
	return nil
}

func (cs *InforceService) AddDept(
	r *http.Request, req *AddDeptReq, resp *AddResp) error {
	c := endpoints.NewContext(r)

	company := &Company{Name: req.CompanyName}
	dept := &Dept{Name: req.DepartmentName, CompanyKey: company.Key(c)}

	_, err := company.AddDept(c, dept)

	if nil != err {
		resp.IsAdded = false
		return err
	}

	resp.IsAdded = true
	return nil
}

func (cs *InforceService) AddUserToCompany(r *http.Request, req *AddUserToCompReq, resp *AddUserToCompResp) error {

	c := endpoints.NewContext(r)
	u, err := GetUser(c, req.Username)
	if nil != err {
		resp.ErrorMsg = "No such user"
		return err
	}

	// Get company
	comp, err := GetCompany(c, req.SelectedCompany)
	if nil != err {
		resp.ErrorMsg = "No such company"
		return err
	}

	userKey := u.Key(c)
	for _, uk := range comp.Users {
		if uk == userKey {
			resp.Success = true
			return nil
		}
	}

	_, err = comp.AddUser(c, u)
	if nil != err {
		resp.ErrorMsg = "Error adding user"
		return err
	}

	resp.Success = true
	return nil
}

func (cs *InforceService) CompanyUsersList(r *http.Request, req *CompanyUsersReq, resp *CompanyUsersResp) error {
	c := endpoints.NewContext(r)

	comp, err := GetCompany(c, req.SelectedCompany)
	if nil != err {
		resp.ErrorMsg = "Company not found"
		return err
	}

	users, err := GetUsers(c, comp.Users)
	if nil != err {
		resp.ErrorMsg = "An error occurred"
		return err
	}

	resp.Users = users
	return nil
}
