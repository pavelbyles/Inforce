package inforce

type IndexVm struct {
	Title        string
	Header       string
	CurrentDate  string
	Companies    *[]Company
	HasCompanies bool
	NumCompanies int
	XsrfToken    string
}

type CompaniesVm struct {
	Companies       *[]Company
	SelectedCompany string
	HasCompanies    bool
	NumCompanies    int
	XsrfToken       string
}

type CompanyVm struct {
	Name             string
	Departments      *[]Dept
	Groups           *[]Group
	IsNoError        bool
	XsrfToken        string
	NumDepartments   int
	NumGroups        int
	CompanyUsers     []UserInfo
	IsUserSessionSet bool
	SessionUser      string
	CompanyMessages  []Message
	HasMessages      bool
	MessageCount     int
}

type AddDeptVm struct {
	CompanyName string
	NewDeptName string
}

type AddCompUserVm struct {
	CompanyName  string
	NewUsername  string
	NewUserFname string
	NewUserLname string
}
