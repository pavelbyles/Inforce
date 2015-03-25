package inforce

import (
	"strconv"
	"testing"

	"appengine/aetest"
)

func Test_numberOf(t *testing.T) {
	company := &Company{
		Name: c1name,
	}

	c, _ := aetest.NewContext(nil)
	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
	}

	numCompanies, err := NumberOf(c, "Company")
	if nil != err {
		t.Errorf("Could not retrieve the number companies: %v", err)
	}

	if 1 != numCompanies {
		t.Errorf("Wrong number of companies. Expected %v, got %v", 1, numCompanies)
	}

	d := &Dept{Name: c1DeptName, CompanyKey: ck}
	_, err = company.AddDept(c, d)
	if nil != err {
		t.Errorf("Could not add department: %v", err)
	}

	numDepts, err := NumberOf(c, "DeptsFor_"+company.Name)
	if nil != err {
		t.Errorf("Could not retrieve the number companies: %v", err)
	}

	if 1 != numDepts {
		t.Errorf("Wrong number of depts.Expected %v, got %v", 1, numDepts)
	}
}

func Test_getAllCompanies(t *testing.T) {
	company := &Company{
		Name: c1name,
	}

	c, _ := aetest.NewContext(nil)
	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}
	companies, err := GetAllCompanies(c)

	if err != nil {
		t.Errorf("Error getting companies: %v", err)
		return
	}

	numComps := len(companies)
	if numComps != 1 {
		t.Errorf("Incorrect number of companies. Expected %v, got %v", 1, numComps)
	}
}

func Test_getAllCompaniesN(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	n := 10

	for i := n; i != 0; i-- {
		comp := &Company{
			Name: strconv.Itoa(i),
		}
		if _, err := AddCompany(c, comp); nil != err {
			t.Errorf("Could not add company: %v", err)
			return
		}
	}

	companies, err := GetAllCompanies(c)
	if err != nil {
		t.Errorf("Could not get all companies: %v", err)
		return
	}

	numComps := len(companies)
	if numComps != n {
		t.Errorf("Incorrect number of companies. Expected %v, got %v", n, numComps)
	}
}

func Test_getCompany(t *testing.T) {
	company := &Company{
		Name: c1name,
	}

	c, _ := aetest.NewContext(nil)
	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not get add company: %v", err)
	}

	companies, err := GetAllCompanies(c)

	if nil != err {
		t.Errorf("Could not get companies: %v", err)
	}

	if len(companies) != 1 {
		t.Errorf("Incorrect number of companies received. Expected %v, got %v", 1, len(companies))
	}

	retComp, err := GetCompany(c, companies[0].Name)
	if nil != err {
		t.Errorf("Could not get company (%v): %v", companies[0].Name, err)
	}

	if retComp.Name != c1name {
		t.Errorf("Expected %v, got %v", c1name, company.Name)
	}
}
