package inforce

import (
	"testing"
	"time"

	"appengine/aetest"
)

func Test_groupMessages(t *testing.T) {
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
	dk, err := company.AddDept(c, dept)
	if nil != err {
		t.Errorf("Could not add dept: %v", err)
		return
	}

	group := &Group{
		Name:    c1GrpName,
		DeptKey: dk,
	}

	if _, err = dept.AddGroup(c, group); nil != err {
		t.Errorf("Could not add group: %v", err)
		return
	}

	m := &Message{Text: "Test Message"}
	var n int = 1
	sTime := time.Now()
	for i := n; i > 0; i-- {
		if _, err = group.AddMessage(c, m); nil != err {
			t.Errorf("Could not add group message: %v", err)
			return
		}
	}
	eTime := time.Now()
	messages, err := group.GetMessages(c, sTime, eTime)

	if nil != err {
		t.Errorf("Could not get messages: %v", err)
		return
	}

	if len(messages) != n {
		t.Errorf("Incorrect number of messages. Expected %d, got %d.", n, len(messages))
		return
	}
}
