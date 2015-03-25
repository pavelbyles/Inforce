package inforce

import (
	"errors"
	"time"

	"appengine"
	"appengine/datastore"
)

type Group struct {
	Name       string
	StartDate  time.Time
	CompanyKey *datastore.Key `datastore:"-"`
	DeptKey    *datastore.Key `datastore:"-"`
}

// Gets a group key
func (group *Group) Key(c appengine.Context) *datastore.Key {

	if nil == group.CompanyKey && nil == group.DeptKey {
		return nil
	}

	if nil != group.CompanyKey {
		return datastore.NewKey(c, "Group", group.Name, 0, group.CompanyKey)
	}
	if nil != group.DeptKey {
		return datastore.NewKey(c, "Group", group.Name, 0, group.DeptKey)
	}
	return nil
}

func (grp *Group) AddMessage(c appengine.Context, m *Message) (*datastore.Key, error) {
	m.Timestamp = time.Now()
	mk := datastore.NewIncompleteKey(c, "Message", grp.Key(c))
	return datastore.Put(c, mk, m)
}

func (grp *Group) GetMessagesByCount(c appengine.Context, lastN int) ([]Message, error) {
	var msgs []Message

	q := datastore.
		NewQuery("Message").Ancestor(grp.Key(c)).
		Order("-Timestamp").
		Limit(lastN)

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}

func (group *Group) GetMessages(c appengine.Context, sTime time.Time, eTime time.Time) ([]Message, error) {
	var msgs []Message

	if sTime.After(eTime) {
		return nil, errors.New("Start time > end time")
	}

	q := datastore.
		NewQuery("Message").Ancestor(group.Key(c)).
		Filter("Timestamp >=", sTime).
		Filter("Timestamp <=", eTime).
		Order("Timestamp")

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}
