package inforce

import (
	"strconv"
	"testing"
	"time"

	"appengine/aetest"
)

func TestAddCompany(t *testing.T) {
	c, _ := aetest.NewContext(nil)

	company := &Company{
		Name: c1name,
	}

	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not add company: %v", err)
	}
}

func TestGetCompany(t *testing.T) {
	c, _ := aetest.NewContext(nil)

	company := &Company{
		Name: c1name,
	}

	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not add company: %v", err)
	}

	comp, err := GetCompany(c, c1name)
	if nil != err {
		t.Errorf("Could not get company: %v", err)
	}

	if !comp.Key(c).Equal(company.Key(c)) {
		t.Errorf("Error in company received. Expected %v, got %v.", company.Key(c), comp.Key(c))
	}

	if comp.Name != company.Name {
		t.Errorf("Error in company received. Expected %v, got %v.", company.Name, comp.Name)
	}
}

func TestGetAllCompanies(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	n := 100
	//var compArr []int

	for i := 0; i < n; i++ {
		company := &Company{
			Name: c1name + strconv.Itoa(i),
		}
		_, err := AddCompany(c, company)
		if nil != err {
			t.Errorf("Could not add company %v: %v", company.Name, err)
		}
	}

	compArr, err := GetAllCompanies(c)
	if nil != err {
		t.Errorf("Could not get companies: %v", err)
	}

	if n != len(compArr) {
		t.Errorf("Incorrect number of companies added. Expected %v, got %v", n, len(compArr))
	}

}

func Test_addDepartment(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{
		Name: c1name,
	}

	ck, err := AddCompany(c, company)

	if nil != err {
		t.Errorf("Could not add company: %v", err)
	}

	dept := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}
	if _, err = company.AddDept(c, dept); nil != err {
		t.Errorf("Could not add dept: %v", err)
	}
}

func Test_addCompGroup(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	group := &Group{Name: c1GrpName, CompanyKey: ck}
	if _, err = company.AddGroup(c, group); nil != err {
		t.Errorf("Could not add group: %v", err)
	}
}

func Test_getAllDepts(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{
		Name: c1name,
	}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	dept := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}
	if _, err = company.AddDept(c, dept); nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}

	depts, err := company.GetDepts(c)

	if nil != err {
		t.Errorf("Could not get departments: %v", err)
		return
	}

	if len(depts) != 1 {
		t.Errorf("Incorrect number of departments. Expected 1, got %d", len(depts))
		return
	}

	if c1DeptName != depts[0].Name {
		t.Errorf("Wrong dept. Expected %s, got %s", c1DeptName, depts[0].Name)
		return
	}
}

func Test_getDept(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{
		Name: c1name,
	}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	dept := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}
	if _, err = company.AddDept(c, dept); nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}

	d, err := company.GetDept(c, dept.Name)
	if nil != err {
		t.Errorf("Could not get dept: %v", err)
	}

	if d.Name != dept.Name {
		t.Errorf("Incorrect dept received. Expected %v, got %v", dept, d)
	}
}

func TestGetCompGroup(t *testing.T) {
	c, _ := aetest.NewContext(nil)

	company := &Company{Name: c1name}
	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	group := &Group{Name: c1GrpName, CompanyKey: ck}
	if _, err = company.AddGroup(c, group); nil != err {
		t.Errorf("Could not add group: %v", err)
		return
	}

	g, err := company.GetGroup(c, c1GrpName)
	if nil != err {
		t.Errorf("Could not get company group: %v", err)
		return
	}
	if c1GrpName != g.Name {
		t.Errorf("Incorrect group returned. Expected %v, got %v", c1GrpName, g.Name)
	}
}

func TestCompanyMessages(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}
	var n int = 100

	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not add company: %v", err)
	}

	u := UserInfo{Username: "wikid", Firstname: "Pavel"}
	m := &Message{Text: "Test Message", FromUser: u}
	sTime := time.Now()
	for i := n; i > 0; i-- {
		if _, err := company.AddMessage(c, m); nil != err {
			t.Errorf("Could not add message to datastore: %v", err)
			return
		}
	}
	eTime := time.Now()

	messages, err := company.GetMessages(c, sTime, eTime)
	if nil != err {
		t.Errorf("Could not get company messages: %v", err)
		return
	}

	if len(messages) != n {
		t.Errorf("Wrong number of messages returned. Expected %d, got %d", n, len(messages))
		return
	}
}

func TestAddCompUser(t *testing.T) {
	c, _ := aetest.NewContext(nil)

	company := &Company{
		Name: c1name,
	}

	user := &UserInfo{
		Username: c1User,
	}

	// Add company
	if _, err := AddCompany(c, company); nil != err {
		t.Errorf("Could not add company: %v", err)
	}
	// Add user
	if _, err := AddUser(c, user); nil != err {
		t.Errorf("Could not add user: %v", err)
	}

	if _, err := company.AddUser(c, user); nil != err {
		t.Errorf("Could not assign user to company: %v", err)
	}
	users, err := company.GetUsers(c)
	if nil != err {
		t.Errorf("Could not get users for company: %v", err)
	}
	if 1 != len(users) {
		t.Errorf("Incorrect number of users assigned to company. Expected %d, got %v", 1, len(users))
	}

	companies, err := user.GetCompanies(c)
	if nil != err {
		t.Errorf("Could not get user companies: %v", err)
	}
	if 1 != len(companies) {
		t.Errorf("Incorrect number of companies assigned to user. Expected %d, got %d", 1, len(companies))
	}
}
