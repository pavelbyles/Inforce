package inforce

import (
	"appengine"
	"appengine/datastore"
	"time"
)

type UserInfo struct {
	Username  string
	Firstname string
	Lastname  string
	TimeAdded time.Time
	Companies []*datastore.Key
}

func (u *UserInfo) Key(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "UserInfo", u.Username, 0, nil)
}

func AddUser(c appengine.Context, user *UserInfo) (*datastore.Key, error) {
	user.TimeAdded = time.Now()
	return datastore.Put(c, user.Key(c), user)
}

func GetUser(c appengine.Context, username string) (*UserInfo, error) {
	user := &UserInfo{Username: username}
	var u UserInfo
	if err := datastore.Get(c, user.Key(c), &u); nil != err {
		return nil, err
	}
	return &u, nil
}

func GetUsers(c appengine.Context, userKeys []*datastore.Key) ([]UserInfo, error) {
	users := make([]UserInfo, len(userKeys))
	if err := datastore.GetMulti(c, userKeys, users); nil != err {
		return nil, err
	}

	return users, nil
}

func (u *UserInfo) GetCompanies(c appengine.Context) ([]Company, error) {
	companies := make([]Company, len(u.Companies))

	if err := datastore.GetMulti(c, u.Companies, companies); nil != err {
		return nil, err
	}

	return companies, nil
}

func (u *UserInfo) AddCompany(c appengine.Context, company *Company) (*datastore.Key, error) {
	u.Companies = append(u.Companies, company.Key(c))
	return datastore.Put(c, u.Key(c), u)
}
