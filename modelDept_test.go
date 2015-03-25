package inforce

import (
	"testing"
	"time"

	"appengine/aetest"
)

func TestAddDeptGroup(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	department := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}
	dk, err := company.AddDept(c, department)
	if nil != err {
		t.Errorf("Could not add department: %v", err)
		return
	}

	group := &Group{Name: c1GrpName, DeptKey: dk}
	if _, err = department.AddGroup(c, group); nil != err {
		t.Errorf("Could not add group: %v", err)
	}
}

func TestGetDeptGroup(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	department := &Dept{Name: c1DeptName, CompanyKey: ck}
	dk, err := company.AddDept(c, department)
	if nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}

	group := &Group{Name: c1GrpName, DeptKey: dk}
	if _, err := department.AddGroup(c, group); nil != err {
		t.Errorf("Could not add group: %v", err)
		return
	}

	grps, err := department.GetGroups(c)
	if nil != err {
		t.Errorf("Could not get groups: %v", err)
	}

	c.Infof("Groups: %v", grps)

	g, err := department.GetGroup(c, c1GrpName)
	if nil != err {
		t.Errorf("Could not get group: %v", err)
		return
	}
	if c1GrpName != g.Name {
		t.Errorf("Incorrect group returned. Expected %v, got %v", c1GrpName, g.Name)
		return
	}
}

func TestDeptMessages(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	dept := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}

	_, err = company.AddDept(c, dept)
	if nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}
	m := &Message{Text: "Test Message", Timestamp: time.Now()}

	var n int = 100
	sTime := time.Now()
	for i := n; i > 0; i-- {
		if _, err = dept.AddMessage(c, m); nil != err {
			t.Errorf("Could not add message: %v", err)
			return
		}
	}
	eTime := time.Now()

	messages, err := dept.GetMessages(c, sTime, eTime)
	if nil != err {
		t.Errorf("Could not get messages: %v", err)
		return
	}

	if len(messages) != n {
		t.Errorf("Incorrect number of messages. Expected %d, got %d.", n, len(messages))
		return
	}
}

func TestGettLastnDeptMessages(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	company := &Company{Name: c1name}

	ck, err := AddCompany(c, company)
	if nil != err {
		t.Errorf("Could not add company: %v", err)
		return
	}

	dept := &Dept{
		Name:       c1DeptName,
		CompanyKey: ck,
	}

	_, err = company.AddDept(c, dept)
	if nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}

	var n int = 100
	for i := n; i > 0; i-- {
		m := &Message{Text: "Test Message", Timestamp: time.Now()}
		if _, err = dept.AddMessage(c, m); nil != err {
			t.Errorf("Could not add message: %v", err)
			return
		}
	}

	messages, err := dept.GetLastnMessages(c, 100)
	if nil != err {
		t.Errorf("Could retieve messages: %v", err)
	}

	if 100 != len(messages) {
		t.Errorf("Could not retrieve all messages. Expected %v, got %v", 100, len(messages))
	}
}
