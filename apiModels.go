package inforce

type LoginReq struct {
	Username string `json:"username"`
	Pasword  string `json:"password"`
}

type LoginResp struct {
	Success   bool      `json:"success"`
	ErrorMsg  string    `json:"errormsg"`
	Companies []Company `json:"companies"`
}

type CompaniesList struct {
	Companies []Company `json:"companies"`
}

type CompaniesListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

type AddCompanyReq struct {
	Name string `json:"name"`
}

type AddDeptReq struct {
	CompanyName    string `json:"companyname"`
	DepartmentName string `json:"deptname"`
}

type AddUserReq struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type AddResp struct {
	IsAdded bool `json:"isadded"`
}

type GroupsList struct {
	Groups []Group `json:"groups"`
}

type GroupsListReq struct {
	SelectedCompany string `json:"companyname"`
	SelectedDept    string `json:"deptname"`
}

type AddUserToCompReq struct {
	Username        string `json:"username"`
	SelectedCompany string `json:"selectedcompany"`
}

type AddUserToCompResp struct {
	Success  bool   `json:"success"`
	ErrorMsg string `json:"errormsg"`
}

type CompaniesForUserReq struct {
	Username string `json:"username"`
}

type CompaniesForUserResp struct {
	//Companies []string `json:"companies"`
	Companies []Company `json:"companies"`
}

type GetUserInfoReq struct {
	Username string `json:"username"`
}

type GetUserInfoResp struct {
	Info     *UserInfo `json:"userinfo"`
	ErrorMsg string    `json:"errormsg"`
}

type CompanyUsersReq struct {
	SelectedCompany string `json:"selectedcompany"`
}

type CompanyUsersResp struct {
	Users    []UserInfo `json:"users"`
	ErrorMsg string     `json:"errormsg"`
}

type UserCompaniesReq struct {
	Username string `json:"username"`
}

type UserCompaniesResp struct {
	Companies []string `json:"companies"`
	ErrorMsg  string   `json:"errormsg"`
}
