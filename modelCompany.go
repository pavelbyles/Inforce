package inforce

import (
	"encoding/json"
	"errors"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/memcache"
)

type Company struct {
	Name  string `json:"Name"`
	Users []*datastore.Key
}

// Gets a company key
func (comp *Company) Key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Company", comp.Name, 0, companyKey(c))
}

// Gets a key that represents the ancestor for all companies
func companyKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Company", "root", 0, nil)
}

// Adds a company
func AddCompany(c appengine.Context, comp *Company) (*datastore.Key, error) {
	key, err := datastore.Put(c, comp.Key(c), comp)
	if nil != err {
		return nil, err
	}
	return key, updateCounter(c, "Company")
}

func GetCompany(c appengine.Context, compName string) (*Company, error) {
	comp := &Company{Name: compName}
	var company Company
	if err := datastore.Get(c, comp.Key(c), &company); err != nil {
		return nil, err
	}
	return &company, nil
}

func GetCompanies(c appengine.Context, companyKeys []*datastore.Key) ([]Company, error) {
	companies := make([]Company, len(companyKeys))
	if err := datastore.GetMulti(c, companyKeys, companies); nil != err {
		return nil, err
	}
	return companies, nil
}

func (comp *Company) AddDept(c appengine.Context, dept *Dept) (*datastore.Key, error) {
	key, err := datastore.Put(c, dept.Key(c), dept)
	if nil != err {
		return nil, err
	}
	return key, updateCounter(c, "DeptsFor_"+comp.Name)
}

func (comp *Company) AddUser(c appengine.Context, user *UserInfo) (*datastore.Key, error) {
	user.TimeAdded = time.Now()
	comp.Users = append(comp.Users, user.Key(c))
	var key *datastore.Key

	txOpts := &datastore.TransactionOptions{XG: true}

	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		// Associate user with company
		key, err1 := datastore.Put(c, comp.Key(c), comp)
		if nil != err1 {
			return err1
		}

		// Associate company with user
		user.Companies = append(user.Companies, key)
		_, err1 = datastore.Put(c, user.Key(c), user)
		if nil != err1 {
			c.Infof("Error associating company with user %v", err1)
			return err1
		}
		return nil
	}, txOpts)
	if nil != err {
		return nil, err
	}
	return key, updateCounter(c, "UsersFor_"+comp.Name)
}

func (comp *Company) AddGroup(c appengine.Context, group *Group) (*datastore.Key, error) {
	return datastore.Put(c, group.Key(c), group)
}

func (comp *Company) AddMessage(c appengine.Context, m *Message) (*datastore.Key, error) {
	m.Timestamp = time.Now()
	mk := datastore.NewIncompleteKey(c, "Message", comp.Key(c))
	return datastore.Put(c, mk, m)
}

// Get all the departments within
func (comp *Company) GetDepts(c appengine.Context) ([]Dept, error) {
	var depts []Dept

	q := datastore.
		NewQuery("Dept").Ancestor(comp.Key(c)).
		Order("Name")

	_, err := q.GetAll(c, &depts)

	return depts, err
}

func (comp *Company) GetDept(c appengine.Context, dname string) (Dept, error) {
	d := &Dept{Name: dname, CompanyKey: comp.Key(c)}
	var dept Dept
	err := datastore.Get(c, d.Key(c), &dept)

	return dept, err
}

func (comp *Company) GetUsers(c appengine.Context) ([]UserInfo, error) {
	users := make([]UserInfo, len(comp.Users))

	if err := datastore.GetMulti(c, comp.Users, users); nil != err {
		return nil, err
	}
	return users, nil
}

func (comp *Company) GetMessagesByCount(c appengine.Context, lastN int) ([]Message, error) {
	var msgs []Message

	q := datastore.
		NewQuery("Message").Ancestor(comp.Key(c)).
		Order("-Timestamp").
		Limit(lastN)

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}

func (comp *Company) GetMessages(c appengine.Context, sTime time.Time, eTime time.Time) ([]Message, error) {
	var msgs []Message

	if sTime.After(eTime) {
		return nil, errors.New("Start time > end time")
	}

	q := datastore.
		NewQuery("Message").Ancestor(comp.Key(c)).
		Filter("Timestamp >=", sTime).
		Filter("Timestamp <=", eTime).
		Order("Timestamp")

	_, err := q.GetAll(c, &msgs)
	return msgs, err
}

func (comp *Company) GetGroup(c appengine.Context,
	gName string) (*Group, error) {

	if len(gName) == 0 {
		err := errors.New("No group specified")
		return nil, err
	}
	refGroup := &Group{Name: gName, CompanyKey: comp.Key(c)}
	var group Group
	if err := datastore.Get(c, refGroup.Key(c), &group); nil != err {
		return nil, err
	}
	return &group, nil
}

func (comp *Company) GetGroups(c appengine.Context) ([]Group, error) {
	var groups []Group
	q := datastore.
		NewQuery("Group").Ancestor(comp.Key(c)).
		Order("Name")
	_, err := q.GetAll(c, &groups)

	return groups, err
}

func GetAllCompanies(c appengine.Context) (companies []Company, err error) {
	if cOutStr, err := memcache.Get(c, "cList"); err == memcache.ErrCacheMiss {
		q := datastore.
			NewQuery("Company").Ancestor(companyKey(c)).
			Order("Name")

		var companies []Company
		keys, err := q.GetAll(c, &companies)
		if err != nil {
			return nil, err
		}

		if len(keys) < 1 {
			return nil, nil
		}

		b, err := json.Marshal(companies)
		if nil != err {
			c.Infof("Error - could not marshal companies to JSON: %v", err)
			return nil, err
		}

		donech := make(chan bool, 1)
		defer close(donech)
		go func() {
			memcache.Set(c, &memcache.Item{
				Key:   "cList",
				Value: []byte(b),
			})
			donech <- true
		}()
		select {
		case <-donech:
			return companies, nil
		}
	} else {
		var companies []Company
		json.Unmarshal(cOutStr.Value, &companies)
		return companies, nil
	}
}
