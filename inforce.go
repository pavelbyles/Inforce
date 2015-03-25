package inforce

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
)

var templates *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func init() {
	templates = template.Must(template.
		ParseFiles(
		"views/_layout.tmpl",
		"views/home/index.tmpl",
		"views/home/_companies.tmpl",
		"views/home/_company.tmpl",
		"views/about/index.tmpl"))

	r := mux.NewRouter()

	r.HandleFunc("/", errorHandler(rootHandler)).Methods("GET")
	r.HandleFunc("/addCompany", errorHandler(addCompanyHandler)).Methods("POST")
	r.HandleFunc("/addDepartment", errorHandler(addDepartmentHandler)).Methods("POST")
	r.HandleFunc("/getCompanies", errorHandler(getCompaniesHandler))
	r.HandleFunc("/getCompany", errorHandler(getCompanyHandler))
	r.HandleFunc("/getCompany/{id:[0-9]+}", errorHandler(getCompanyByIdHandler))
	r.HandleFunc("/sendCompanyMsg", errorHandler(sendCompanyMsgHandler))
	r.HandleFunc("/sendDeptMsg", errorHandler(sendDeptMsgHandler))
	r.HandleFunc("/sendGroupMsg", errorHandler(sendGroupMsgHandler))
	r.HandleFunc("/sendP2pMsg", errorHandler(sendP2pMsgHandler))
	r.HandleFunc("/addCompUser", errorHandler(addCompUserHandler))
	r.HandleFunc("/setUser", errorHandler(setUserHandler))

	r.HandleFunc("/about", errorHandler(aboutPageHandler)).Methods("GET")

	http.Handle("/", r)
	registerServices()

	//endpoints.HandleHTTP()
}

// Root handler
func rootHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	t := time.Now()
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	companies, err := GetAllCompanies(c)

	if err != nil {
		return err
	}

	m := &IndexVm{Title: "Test", Header: "PaX Test",
		CurrentDate: t.Format(layout), Companies: &companies,
		HasCompanies: (len(companies) > 0), NumCompanies: (len(companies))}
	return renderIndexTmpl(w, m)
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) error {
	return renderAboutTmpl(w)
}

func addCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	compName := r.FormValue("companyName")

	if len(compName) < 1 {
		w.WriteHeader(200)
		return nil
	}

	company := &Company{
		Name: compName,
	}

	if _, err := AddCompany(c, company); nil != err {
		c.Infof("Error: error adding company: %v", err)
		return err
	}

	memcache.Delete(c, "cList")
	w.WriteHeader(200)
	return nil
}

func setUserHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	username := r.FormValue("username")

	user, err := GetUser(c, username)
	if nil != err {
		c.Infof("No user found:" + err.Error())
		err = errors.New("No user found: " + err.Error())
		return err
	}

	session, _ := store.Get(r, "default")
	session.Values["username"] = user.Username
	session.Values["firstname"] = user.Firstname
	session.Values["lastname"] = user.Lastname

	session.Save(r, w)
	w.WriteHeader(200)
	return nil
}

func addCompUserHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	err := r.ParseForm()

	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	d := schema.NewDecoder()

	acuvm := new(AddCompUserVm)
	if err := d.Decode(acuvm, r.PostForm); nil != err {
		return err
	}

	if len(acuvm.CompanyName) < 1 || len(acuvm.NewUsername) < 1 ||
		len(acuvm.NewUserFname) < 1 || len(acuvm.NewUserLname) < 1 {
		w.WriteHeader(200)
		return nil
	}
	company, err := GetCompany(c, acuvm.CompanyName)
	if nil != err {
		c.Infof("Error: company not found: %v", err)
		return err
	}

	user, err := GetUser(c, acuvm.NewUsername)
	if nil != err {
		user = &UserInfo{
			Username:  acuvm.NewUsername,
			Firstname: acuvm.NewUserFname,
			Lastname:  acuvm.NewUserLname,
		}
	}

	_, err = company.AddUser(c, user)

	if nil != err {
		return err
	}

	w.WriteHeader(200)
	return nil
}

func sendCompanyMsgHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	compName := r.FormValue("company")
	msgToAdd := r.FormValue("companyMessage")
	username := r.FormValue("username")

	if len(compName) < 1 || len(msgToAdd) < 1 {
		w.WriteHeader(200)
		return nil
	}

	user, err := GetUser(c, username)
	if nil != err {
		c.Infof("Error getting user: %v", err)
		return err
	}

	m := &Message{
		FromUser: *user,
		Text:     msgToAdd,
	}

	company, err := GetCompany(c, compName)
	if nil != err {
		c.Infof("Error getting user: %v", err)
		return err
	}

	_, err = company.AddMessage(c, m)
	if nil != err {
		c.Infof("Error adding message: %v", err)
		return err
	}

	if _, err = AddCompany(c, company); nil != err {
		c.Infof("Couldn't update company with user: %v", err)
		return err
	}

	return nil
}

func sendDeptMsgHandler(w http.ResponseWriter, r *http.Request) error {
	//c := appengine.NewContext(r)
	return nil
}

func sendGroupMsgHandler(w http.ResponseWriter, r *http.Request) error {
	//c := appengine.NewContext(r)
	return nil
}

func sendP2pMsgHandler(w http.ResponseWriter, r *http.Request) error {
	//c := appengine.NewContext(r)
	return nil
}

func addDepartmentHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	err := r.ParseForm()

	if nil != err {
		return err
	}

	d := schema.NewDecoder()
	advm := new(AddDeptVm)
	if err := d.Decode(advm, r.PostForm); nil != err {
		return err
	}

	company, err := GetCompany(c, advm.CompanyName)

	dept := &Dept{
		Name:       advm.NewDeptName,
		CompanyKey: company.Key(c),
	}
	if _, err = company.AddDept(c, dept); nil != err {
		c.Infof("Error: error adding department: %v", err)
		return err
	}
	return nil
}

func getCompaniesHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)

	outputType := r.FormValue("type")

	if companies, err := GetAllCompanies(c); err != nil {
		c.Infof("Error - could not get companies: %v", err)
		return err
	} else {
		if "json" == outputType || "" == outputType {

			b, err := json.Marshal(companies)
			if nil != err {
				return err
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(b))
		} else if "html" == outputType {
			hasCs := (len(companies) > 0)
			cvm := &CompaniesVm{
				Companies:       &companies,
				SelectedCompany: "Sagicor",
				HasCompanies:    hasCs,
				NumCompanies:    len(companies),
			}
			rendLstCompsPartTmpl(w, cvm)
		}
		if nil != err {
			c.Infof("Error - could not convert to JSON: %v", err)
			return err
		}
	}
	return nil
}

func getCompanyHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	err := r.ParseForm()

	if nil != err {
		return err
	}

	d := schema.NewDecoder()
	cvm := new(CompaniesVm)
	if err := d.Decode(cvm, r.PostForm); nil != err {
		return err
	}

	company, err := GetCompany(c, cvm.SelectedCompany)
	if nil != err {
		c.Infof("No company found: %v", cvm.SelectedCompany)
		return err
	}
	departments, err := company.GetDepts(c)
	if nil != err {
		return err
	}
	groups, err := company.GetGroups(c)
	if nil != err {
		return err
	}

	session, _ := store.Get(r, "default")
	su := ""
	if len(company.Users) > 0 && nil != session.Values["username"] {
		su = session.Values["username"].(string)
	}

	isUserSessionSet := false
	if _, ok := session.Values["username"]; ok {
		isUserSessionSet = true
	}

	messages, err := company.GetMessagesByCount(c, 10)
	if nil != err {
		return err
	}
	//hasMsg := len(messages) > 0

	companyUsers, err := GetUsers(c, company.Users)
	if nil != err {
		return err
	}

	m := &CompanyVm{
		Name:             company.Name,
		Departments:      &departments,
		Groups:           &groups,
		NumDepartments:   len(departments),
		NumGroups:        len(groups),
		IsNoError:        true,
		CompanyUsers:     companyUsers,
		IsUserSessionSet: isUserSessionSet,
		SessionUser:      su,
		CompanyMessages:  messages,
		HasMessages:      len(messages) > 0,
		MessageCount:     len(messages),
	}

	//c.Infof("model: %v", m)

	return renderCompPartialTmpl(w, m)
}

func getCompanyByIdHandler(w http.ResponseWriter, r *http.Request) error {
	c := appengine.NewContext(r)
	vars := mux.Vars(r)
	companyName := vars["id"]

	company, err := GetCompany(c, companyName)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(company)
}

func renderIndexTmpl(w http.ResponseWriter, m *IndexVm) error {
	return templates.ExecuteTemplate(w, "index", m)
}

func rendLstCompsPartTmpl(w http.ResponseWriter, m *CompaniesVm) string {
	var doc bytes.Buffer
	err := templates.ExecuteTemplate(&doc, "selectcompanies", m)
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return ""
	}
	fmt.Fprint(w, doc.String())
	return doc.String()
}

func renderCompPartialTmpl(w http.ResponseWriter, m *CompanyVm) error {
	return templates.ExecuteTemplate(w, "showcompany", m)
}

func renderAboutTmpl(w http.ResponseWriter) error {
	return templates.ExecuteTemplate(w, "about", nil)
}

func getDeptGroups(c appengine.Context, cName string,
	dName string) ([]Group, error) {

	company := &Company{Name: cName}
	ck := company.Key(c)
	dept := &Dept{Name: dName, CompanyKey: ck}
	var groups []Group
	c.Infof("Department Key: %v", dept.Key(c))

	q := datastore.
		NewQuery("Group").Ancestor(dept.Key(c)).
		Order("Name")
	_, err := q.GetAll(c, &groups)

	return groups, err
}

func getMessageKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Message", "default_message", 0, nil)
}
