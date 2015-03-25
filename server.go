package inforce

import (
	"github.com/crhym3/go-endpoints/endpoints"

	"log"
	"net/http"
)

type badRequest struct{ error }
type notFound struct{ error }

func errorHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "task not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}

}

func svcErrorHandler(api *endpoints.RPCService, apierr error) *endpoints.RPCService {
	if nil != apierr {
		panic(apierr.Error())
	}

	return api
}

func registerServices() {
	svc := &InforceService{}
	api := svcErrorHandler(endpoints.RegisterService(svc,
		"inforce", "v1", "Companies API", true))

	info := api.MethodByName("CompanyList").Info()
	info.Name, info.HTTPMethod, info.Path, info.Desc = "companies.list", "GET",
		"companies", "List all companies."

	gInfo := api.MethodByName("DeptGroups").Info()
	gInfo.Name, gInfo.HTTPMethod, gInfo.Path, gInfo.Desc = "departments.groups",
		"GET", "groups", "List all groups in a specified company's department."

	cInfo := api.MethodByName("AddCompany").Info()
	cInfo.Name, cInfo.HTTPMethod, cInfo.Path, cInfo.Desc = "companies.add", "POST", "addCompanyResult", "Adds a company."

	addDeptInfo := api.MethodByName("AddDept").Info()
	addDeptInfo.Name, addDeptInfo.HTTPMethod, addDeptInfo.Path, addDeptInfo.Desc = "departments.add", "POST", "addDeptResult", "Adds a department."

	valUserInfo := api.MethodByName("ValidateUserGetCompanies").Info()
	valUserInfo.Name, valUserInfo.HTTPMethod, valUserInfo.Path, valUserInfo.Desc = "userInfos.validate", "POST", "valUserInfoResult", "Validates user and returns the companies belonged to."

	addUserToCompInfo := api.MethodByName("AddUserToCompany").Info()
	addUserToCompInfo.Name, addUserToCompInfo.HTTPMethod, addUserToCompInfo.Path, addUserToCompInfo.Desc = "companies.adduser", "POST", "addUserToCompResult", "Adds a user to a company."

	compUsersInfo := api.MethodByName("CompanyUsersList").Info()
	compUsersInfo.Name, compUsersInfo.HTTPMethod, compUsersInfo.Path, compUsersInfo.Desc = "companies.userlist", "GET", "usersResult", "Gets users within a company."

	userCompInfo := api.MethodByName("UserCompanyList").Info()
	userCompInfo.Name, userCompInfo.HTTPMethod, userCompInfo.Path, userCompInfo.Desc = "userInfos.companylist", "GET", "userCompaniesResult", "Get companies that a user is in."

	getUserInfo := api.MethodByName("GetUserInfo").Info()
	getUserInfo.Name, getUserInfo.HTTPMethod, getUserInfo.Path, getUserInfo.Desc = "userInfos.info", "GET", "userResult", "Get user."

	addUserInfo := api.MethodByName("AddUser").Info()
	addUserInfo.Name, addUserInfo.HTTPMethod, addUserInfo.Path, addUserInfo.Desc = "userInfos.add", "POST", "addUserResult", "Add user."
	endpoints.HandleHTTP()
}
