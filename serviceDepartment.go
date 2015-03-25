package inforce

import (
	"errors"
	"github.com/crhym3/go-endpoints/endpoints"
	"net/http"
)

func (cs *InforceService) DeptGroups(
	r *http.Request, req *GroupsListReq, resp *GroupsList) error {

	c := endpoints.NewContext(r)

	if "" == req.SelectedCompany {
		err := errors.New("No company selected")
		return err
	}

	if "" == req.SelectedDept {
		err := errors.New("No department selected")
		return err
	}

	company := &Company{Name: req.SelectedCompany}
	dept := &Dept{Name: req.SelectedDept, CompanyKey: company.Key(c)}
	grps, err := dept.GetGroups(c)

	if nil != err {
		return err
	}

	resp.Groups = grps
	return nil
}
