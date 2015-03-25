package inforce

import (
	"errors"
	"time"

	"appengine"
	"appengine/datastore"
)

type Dept struct {
	Name       string
	CompanyKey *datastore.Key `datastore:"-"`
}

// Gets a department key
func (d *Dept) Key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Dept", d.Name, 0, d.CompanyKey)
}

func (dept *Dept) AddGroup(c appengine.Context, group *Group) (*datastore.Key, error) {
	return datastore.Put(c, group.Key(c), group)
}

func (dept *Dept) AddMessage(c appengine.Context, m *Message) (*datastore.Key, error) {
	m.Timestamp = time.Now()
	mk := datastore.NewIncompleteKey(c, "Message", dept.Key(c))
	return datastore.Put(c, mk, m)
}

func (dept *Dept) GetGroups(c appengine.Context) ([]Group, error) {
	var grps []Group

	q := datastore.
		NewQuery("Group").Ancestor(dept.Key(c))

	_, err := q.GetAll(c, &grps)
	return grps, err
}

func (dept *Dept) GetGroup(c appengine.Context, gName string) (*Group, error) {
	refGroup := &Group{Name: gName, DeptKey: dept.Key(c)}
	var g Group
	if err := datastore.Get(c, refGroup.Key(c), &g); nil != err {
		return nil, err
	}
	return &g, nil
}

func (dept *Dept) GetLastnMessages(c appengine.Context, lastN int) ([]Message, error) {
	var msgs []Message

	q := datastore.
		NewQuery("Message").Ancestor(dept.Key(c)).
		Order("-Timestamp").
		Limit(lastN)

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}

func (dept *Dept) GetMessages(c appengine.Context, sTime time.Time, eTime time.Time) ([]Message, error) {
	if sTime.After(eTime) {
		return nil, errors.New("Start time > end time")
	}

	var msgs []Message

	q := datastore.
		NewQuery("Message").Ancestor(dept.Key(c)).
		Filter("Timestamp >=", sTime).
		Filter("Timestamp <=", eTime).
		Order("Timestamp")

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}
