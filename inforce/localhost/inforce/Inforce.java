/*
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */
/*
 * This code was generated by https://code.google.com/p/google-apis-client-generator/
 * (build: 2014-11-17 18:43:33 UTC)
 * on 2015-01-05 at 04:45:47 UTC 
 * Modify at your own risk.
 */

package localhost.inforce;

/**
 * Service definition for Inforce (v1).
 *
 * <p>
 * Companies API
 * </p>
 *
 * <p>
 * For more information about this service, see the
 * <a href="" target="_blank">API Documentation</a>
 * </p>
 *
 * <p>
 * This service uses {@link InforceRequestInitializer} to initialize global parameters via its
 * {@link Builder}.
 * </p>
 *
 * @since 1.3
 * @author Google, Inc.
 */
@SuppressWarnings("javadoc")
public class Inforce extends com.google.api.client.googleapis.services.json.AbstractGoogleJsonClient {

  // Note: Leave this static initializer at the top of the file.
  static {
    com.google.api.client.util.Preconditions.checkState(
        com.google.api.client.googleapis.GoogleUtils.MAJOR_VERSION == 1 &&
        com.google.api.client.googleapis.GoogleUtils.MINOR_VERSION >= 15,
        "You are currently running with version %s of google-api-client. " +
        "You need at least version 1.15 of google-api-client to run version " +
        "1.19.0 of the inforce library.", com.google.api.client.googleapis.GoogleUtils.VERSION);
  }

  /**
   * The default encoded root URL of the service. This is determined when the library is generated
   * and normally should not be changed.
   *
   * @since 1.7
   */
  public static final String DEFAULT_ROOT_URL = "http://localhost:8080/_ah/api/";

  /**
   * The default encoded service path of the service. This is determined when the library is
   * generated and normally should not be changed.
   *
   * @since 1.7
   */
  public static final String DEFAULT_SERVICE_PATH = "inforce/v1/";

  /**
   * The default encoded base URL of the service. This is determined when the library is generated
   * and normally should not be changed.
   */
  public static final String DEFAULT_BASE_URL = DEFAULT_ROOT_URL + DEFAULT_SERVICE_PATH;

  /**
   * Constructor.
   *
   * <p>
   * Use {@link Builder} if you need to specify any of the optional parameters.
   * </p>
   *
   * @param transport HTTP transport, which should normally be:
   *        <ul>
   *        <li>Google App Engine:
   *        {@code com.google.api.client.extensions.appengine.http.UrlFetchTransport}</li>
   *        <li>Android: {@code newCompatibleTransport} from
   *        {@code com.google.api.client.extensions.android.http.AndroidHttp}</li>
   *        <li>Java: {@link com.google.api.client.googleapis.javanet.GoogleNetHttpTransport#newTrustedTransport()}
   *        </li>
   *        </ul>
   * @param jsonFactory JSON factory, which may be:
   *        <ul>
   *        <li>Jackson: {@code com.google.api.client.json.jackson2.JacksonFactory}</li>
   *        <li>Google GSON: {@code com.google.api.client.json.gson.GsonFactory}</li>
   *        <li>Android Honeycomb or higher:
   *        {@code com.google.api.client.extensions.android.json.AndroidJsonFactory}</li>
   *        </ul>
   * @param httpRequestInitializer HTTP request initializer or {@code null} for none
   * @since 1.7
   */
  public Inforce(com.google.api.client.http.HttpTransport transport, com.google.api.client.json.JsonFactory jsonFactory,
      com.google.api.client.http.HttpRequestInitializer httpRequestInitializer) {
    this(new Builder(transport, jsonFactory, httpRequestInitializer));
  }

  /**
   * @param builder builder
   */
  Inforce(Builder builder) {
    super(builder);
  }

  @Override
  protected void initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest<?> httpClientRequest) throws java.io.IOException {
    super.initialize(httpClientRequest);
  }

  /**
   * An accessor for creating requests from the Companies collection.
   *
   * <p>The typical use is:</p>
   * <pre>
   *   {@code Inforce inforce = new Inforce(...);}
   *   {@code Inforce.Companies.List request = inforce.companies().list(parameters ...)}
   * </pre>
   *
   * @return the resource collection
   */
  public Companies companies() {
    return new Companies();
  }

  /**
   * The "companies" collection of methods.
   */
  public class Companies {

    /**
     * Adds a company.
     *
     * Create a request for the method "companies.add".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Add#execute()} method to invoke the remote operation.
     *
     * @param content the {@link localhost.inforce.model.AddCompanyReq}
     * @return the request
     */
    public Add add(localhost.inforce.model.AddCompanyReq content) throws java.io.IOException {
      Add result = new Add(content);
      initialize(result);
      return result;
    }

    public class Add extends InforceRequest<localhost.inforce.model.AddResp> {

      private static final String REST_PATH = "addCompanyResult";

      /**
       * Adds a company.
       *
       * Create a request for the method "companies.add".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Add#execute()} method to invoke the remote operation. <p> {@link
       * Add#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must be
       * called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @param content the {@link localhost.inforce.model.AddCompanyReq}
       * @since 1.13
       */
      protected Add(localhost.inforce.model.AddCompanyReq content) {
        super(Inforce.this, "POST", REST_PATH, content, localhost.inforce.model.AddResp.class);
      }

      @Override
      public Add setAlt(java.lang.String alt) {
        return (Add) super.setAlt(alt);
      }

      @Override
      public Add setFields(java.lang.String fields) {
        return (Add) super.setFields(fields);
      }

      @Override
      public Add setKey(java.lang.String key) {
        return (Add) super.setKey(key);
      }

      @Override
      public Add setOauthToken(java.lang.String oauthToken) {
        return (Add) super.setOauthToken(oauthToken);
      }

      @Override
      public Add setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Add) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Add setQuotaUser(java.lang.String quotaUser) {
        return (Add) super.setQuotaUser(quotaUser);
      }

      @Override
      public Add setUserIp(java.lang.String userIp) {
        return (Add) super.setUserIp(userIp);
      }

      @Override
      public Add set(String parameterName, Object value) {
        return (Add) super.set(parameterName, value);
      }
    }
    /**
     * Adds a user to a company.
     *
     * Create a request for the method "companies.adduser".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Adduser#execute()} method to invoke the remote operation.
     *
     * @param content the {@link localhost.inforce.model.AddUserToCompReq}
     * @return the request
     */
    public Adduser adduser(localhost.inforce.model.AddUserToCompReq content) throws java.io.IOException {
      Adduser result = new Adduser(content);
      initialize(result);
      return result;
    }

    public class Adduser extends InforceRequest<localhost.inforce.model.AddUserToCompResp> {

      private static final String REST_PATH = "addUserToCompResult";

      /**
       * Adds a user to a company.
       *
       * Create a request for the method "companies.adduser".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Adduser#execute()} method to invoke the remote operation. <p>
       * {@link
       * Adduser#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must
       * be called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @param content the {@link localhost.inforce.model.AddUserToCompReq}
       * @since 1.13
       */
      protected Adduser(localhost.inforce.model.AddUserToCompReq content) {
        super(Inforce.this, "POST", REST_PATH, content, localhost.inforce.model.AddUserToCompResp.class);
      }

      @Override
      public Adduser setAlt(java.lang.String alt) {
        return (Adduser) super.setAlt(alt);
      }

      @Override
      public Adduser setFields(java.lang.String fields) {
        return (Adduser) super.setFields(fields);
      }

      @Override
      public Adduser setKey(java.lang.String key) {
        return (Adduser) super.setKey(key);
      }

      @Override
      public Adduser setOauthToken(java.lang.String oauthToken) {
        return (Adduser) super.setOauthToken(oauthToken);
      }

      @Override
      public Adduser setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Adduser) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Adduser setQuotaUser(java.lang.String quotaUser) {
        return (Adduser) super.setQuotaUser(quotaUser);
      }

      @Override
      public Adduser setUserIp(java.lang.String userIp) {
        return (Adduser) super.setUserIp(userIp);
      }

      @Override
      public Adduser set(String parameterName, Object value) {
        return (Adduser) super.set(parameterName, value);
      }
    }
    /**
     * List all companies.
     *
     * Create a request for the method "companies.list".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link List#execute()} method to invoke the remote operation.
     *
     * @return the request
     */
    public List list() throws java.io.IOException {
      List result = new List();
      initialize(result);
      return result;
    }

    public class List extends InforceRequest<localhost.inforce.model.CompaniesList> {

      private static final String REST_PATH = "companies";

      /**
       * List all companies.
       *
       * Create a request for the method "companies.list".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link List#execute()} method to invoke the remote operation. <p> {@link
       * List#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must be
       * called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @since 1.13
       */
      protected List() {
        super(Inforce.this, "GET", REST_PATH, null, localhost.inforce.model.CompaniesList.class);
      }

      @Override
      public com.google.api.client.http.HttpResponse executeUsingHead() throws java.io.IOException {
        return super.executeUsingHead();
      }

      @Override
      public com.google.api.client.http.HttpRequest buildHttpRequestUsingHead() throws java.io.IOException {
        return super.buildHttpRequestUsingHead();
      }

      @Override
      public List setAlt(java.lang.String alt) {
        return (List) super.setAlt(alt);
      }

      @Override
      public List setFields(java.lang.String fields) {
        return (List) super.setFields(fields);
      }

      @Override
      public List setKey(java.lang.String key) {
        return (List) super.setKey(key);
      }

      @Override
      public List setOauthToken(java.lang.String oauthToken) {
        return (List) super.setOauthToken(oauthToken);
      }

      @Override
      public List setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (List) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public List setQuotaUser(java.lang.String quotaUser) {
        return (List) super.setQuotaUser(quotaUser);
      }

      @Override
      public List setUserIp(java.lang.String userIp) {
        return (List) super.setUserIp(userIp);
      }

      @com.google.api.client.util.Key
      private java.lang.Integer limit;

      /**
     [ default: 10]
     [

       */
      public java.lang.Integer getLimit() {
        return limit;
      }

      public List setLimit(java.lang.Integer limit) {
        this.limit = limit;
        return this;
      }

      @Override
      public List set(String parameterName, Object value) {
        return (List) super.set(parameterName, value);
      }
    }
    /**
     * Gets users within a company.
     *
     * Create a request for the method "companies.userlist".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Userlist#execute()} method to invoke the remote operation.
     *
     * @return the request
     */
    public Userlist userlist() throws java.io.IOException {
      Userlist result = new Userlist();
      initialize(result);
      return result;
    }

    public class Userlist extends InforceRequest<localhost.inforce.model.CompanyUsersResp> {

      private static final String REST_PATH = "usersResult";

      /**
       * Gets users within a company.
       *
       * Create a request for the method "companies.userlist".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Userlist#execute()} method to invoke the remote operation. <p>
       * {@link
       * Userlist#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)}
       * must be called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @since 1.13
       */
      protected Userlist() {
        super(Inforce.this, "GET", REST_PATH, null, localhost.inforce.model.CompanyUsersResp.class);
      }

      @Override
      public com.google.api.client.http.HttpResponse executeUsingHead() throws java.io.IOException {
        return super.executeUsingHead();
      }

      @Override
      public com.google.api.client.http.HttpRequest buildHttpRequestUsingHead() throws java.io.IOException {
        return super.buildHttpRequestUsingHead();
      }

      @Override
      public Userlist setAlt(java.lang.String alt) {
        return (Userlist) super.setAlt(alt);
      }

      @Override
      public Userlist setFields(java.lang.String fields) {
        return (Userlist) super.setFields(fields);
      }

      @Override
      public Userlist setKey(java.lang.String key) {
        return (Userlist) super.setKey(key);
      }

      @Override
      public Userlist setOauthToken(java.lang.String oauthToken) {
        return (Userlist) super.setOauthToken(oauthToken);
      }

      @Override
      public Userlist setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Userlist) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Userlist setQuotaUser(java.lang.String quotaUser) {
        return (Userlist) super.setQuotaUser(quotaUser);
      }

      @Override
      public Userlist setUserIp(java.lang.String userIp) {
        return (Userlist) super.setUserIp(userIp);
      }

      @com.google.api.client.util.Key
      private java.lang.String selectedcompany;

      /**

       */
      public java.lang.String getSelectedcompany() {
        return selectedcompany;
      }

      public Userlist setSelectedcompany(java.lang.String selectedcompany) {
        this.selectedcompany = selectedcompany;
        return this;
      }

      @Override
      public Userlist set(String parameterName, Object value) {
        return (Userlist) super.set(parameterName, value);
      }
    }

  }

  /**
   * An accessor for creating requests from the Departments collection.
   *
   * <p>The typical use is:</p>
   * <pre>
   *   {@code Inforce inforce = new Inforce(...);}
   *   {@code Inforce.Departments.List request = inforce.departments().list(parameters ...)}
   * </pre>
   *
   * @return the resource collection
   */
  public Departments departments() {
    return new Departments();
  }

  /**
   * The "departments" collection of methods.
   */
  public class Departments {

    /**
     * Adds a department.
     *
     * Create a request for the method "departments.add".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Add#execute()} method to invoke the remote operation.
     *
     * @param content the {@link localhost.inforce.model.AddDeptReq}
     * @return the request
     */
    public Add add(localhost.inforce.model.AddDeptReq content) throws java.io.IOException {
      Add result = new Add(content);
      initialize(result);
      return result;
    }

    public class Add extends InforceRequest<localhost.inforce.model.AddResp> {

      private static final String REST_PATH = "addDeptResult";

      /**
       * Adds a department.
       *
       * Create a request for the method "departments.add".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Add#execute()} method to invoke the remote operation. <p> {@link
       * Add#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must be
       * called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @param content the {@link localhost.inforce.model.AddDeptReq}
       * @since 1.13
       */
      protected Add(localhost.inforce.model.AddDeptReq content) {
        super(Inforce.this, "POST", REST_PATH, content, localhost.inforce.model.AddResp.class);
      }

      @Override
      public Add setAlt(java.lang.String alt) {
        return (Add) super.setAlt(alt);
      }

      @Override
      public Add setFields(java.lang.String fields) {
        return (Add) super.setFields(fields);
      }

      @Override
      public Add setKey(java.lang.String key) {
        return (Add) super.setKey(key);
      }

      @Override
      public Add setOauthToken(java.lang.String oauthToken) {
        return (Add) super.setOauthToken(oauthToken);
      }

      @Override
      public Add setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Add) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Add setQuotaUser(java.lang.String quotaUser) {
        return (Add) super.setQuotaUser(quotaUser);
      }

      @Override
      public Add setUserIp(java.lang.String userIp) {
        return (Add) super.setUserIp(userIp);
      }

      @Override
      public Add set(String parameterName, Object value) {
        return (Add) super.set(parameterName, value);
      }
    }
    /**
     * List all groups in a specified company's department.
     *
     * Create a request for the method "departments.groups".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Groups#execute()} method to invoke the remote operation.
     *
     * @return the request
     */
    public Groups groups() throws java.io.IOException {
      Groups result = new Groups();
      initialize(result);
      return result;
    }

    public class Groups extends InforceRequest<localhost.inforce.model.GroupsList> {

      private static final String REST_PATH = "groups";

      /**
       * List all groups in a specified company's department.
       *
       * Create a request for the method "departments.groups".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Groups#execute()} method to invoke the remote operation. <p> {@link
       * Groups#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must
       * be called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @since 1.13
       */
      protected Groups() {
        super(Inforce.this, "GET", REST_PATH, null, localhost.inforce.model.GroupsList.class);
      }

      @Override
      public com.google.api.client.http.HttpResponse executeUsingHead() throws java.io.IOException {
        return super.executeUsingHead();
      }

      @Override
      public com.google.api.client.http.HttpRequest buildHttpRequestUsingHead() throws java.io.IOException {
        return super.buildHttpRequestUsingHead();
      }

      @Override
      public Groups setAlt(java.lang.String alt) {
        return (Groups) super.setAlt(alt);
      }

      @Override
      public Groups setFields(java.lang.String fields) {
        return (Groups) super.setFields(fields);
      }

      @Override
      public Groups setKey(java.lang.String key) {
        return (Groups) super.setKey(key);
      }

      @Override
      public Groups setOauthToken(java.lang.String oauthToken) {
        return (Groups) super.setOauthToken(oauthToken);
      }

      @Override
      public Groups setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Groups) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Groups setQuotaUser(java.lang.String quotaUser) {
        return (Groups) super.setQuotaUser(quotaUser);
      }

      @Override
      public Groups setUserIp(java.lang.String userIp) {
        return (Groups) super.setUserIp(userIp);
      }

      @com.google.api.client.util.Key
      private java.lang.String deptname;

      /**

       */
      public java.lang.String getDeptname() {
        return deptname;
      }

      public Groups setDeptname(java.lang.String deptname) {
        this.deptname = deptname;
        return this;
      }

      @com.google.api.client.util.Key
      private java.lang.String companyname;

      /**

       */
      public java.lang.String getCompanyname() {
        return companyname;
      }

      public Groups setCompanyname(java.lang.String companyname) {
        this.companyname = companyname;
        return this;
      }

      @Override
      public Groups set(String parameterName, Object value) {
        return (Groups) super.set(parameterName, value);
      }
    }

  }

  /**
   * An accessor for creating requests from the UserInfos collection.
   *
   * <p>The typical use is:</p>
   * <pre>
   *   {@code Inforce inforce = new Inforce(...);}
   *   {@code Inforce.UserInfos.List request = inforce.userInfos().list(parameters ...)}
   * </pre>
   *
   * @return the resource collection
   */
  public UserInfos userInfos() {
    return new UserInfos();
  }

  /**
   * The "userInfos" collection of methods.
   */
  public class UserInfos {

    /**
     * Add user.
     *
     * Create a request for the method "userInfos.add".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Add#execute()} method to invoke the remote operation.
     *
     * @param content the {@link localhost.inforce.model.AddUserReq}
     * @return the request
     */
    public Add add(localhost.inforce.model.AddUserReq content) throws java.io.IOException {
      Add result = new Add(content);
      initialize(result);
      return result;
    }

    public class Add extends InforceRequest<localhost.inforce.model.AddResp> {

      private static final String REST_PATH = "addUserResult";

      /**
       * Add user.
       *
       * Create a request for the method "userInfos.add".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Add#execute()} method to invoke the remote operation. <p> {@link
       * Add#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must be
       * called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @param content the {@link localhost.inforce.model.AddUserReq}
       * @since 1.13
       */
      protected Add(localhost.inforce.model.AddUserReq content) {
        super(Inforce.this, "POST", REST_PATH, content, localhost.inforce.model.AddResp.class);
      }

      @Override
      public Add setAlt(java.lang.String alt) {
        return (Add) super.setAlt(alt);
      }

      @Override
      public Add setFields(java.lang.String fields) {
        return (Add) super.setFields(fields);
      }

      @Override
      public Add setKey(java.lang.String key) {
        return (Add) super.setKey(key);
      }

      @Override
      public Add setOauthToken(java.lang.String oauthToken) {
        return (Add) super.setOauthToken(oauthToken);
      }

      @Override
      public Add setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Add) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Add setQuotaUser(java.lang.String quotaUser) {
        return (Add) super.setQuotaUser(quotaUser);
      }

      @Override
      public Add setUserIp(java.lang.String userIp) {
        return (Add) super.setUserIp(userIp);
      }

      @Override
      public Add set(String parameterName, Object value) {
        return (Add) super.set(parameterName, value);
      }
    }
    /**
     * Get companies that a user is in.
     *
     * Create a request for the method "userInfos.companylist".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Companylist#execute()} method to invoke the remote operation.
     *
     * @return the request
     */
    public Companylist companylist() throws java.io.IOException {
      Companylist result = new Companylist();
      initialize(result);
      return result;
    }

    public class Companylist extends InforceRequest<localhost.inforce.model.UserCompaniesResp> {

      private static final String REST_PATH = "userCompaniesResult";

      /**
       * Get companies that a user is in.
       *
       * Create a request for the method "userInfos.companylist".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Companylist#execute()} method to invoke the remote operation. <p>
       * {@link
       * Companylist#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)}
       * must be called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @since 1.13
       */
      protected Companylist() {
        super(Inforce.this, "GET", REST_PATH, null, localhost.inforce.model.UserCompaniesResp.class);
      }

      @Override
      public com.google.api.client.http.HttpResponse executeUsingHead() throws java.io.IOException {
        return super.executeUsingHead();
      }

      @Override
      public com.google.api.client.http.HttpRequest buildHttpRequestUsingHead() throws java.io.IOException {
        return super.buildHttpRequestUsingHead();
      }

      @Override
      public Companylist setAlt(java.lang.String alt) {
        return (Companylist) super.setAlt(alt);
      }

      @Override
      public Companylist setFields(java.lang.String fields) {
        return (Companylist) super.setFields(fields);
      }

      @Override
      public Companylist setKey(java.lang.String key) {
        return (Companylist) super.setKey(key);
      }

      @Override
      public Companylist setOauthToken(java.lang.String oauthToken) {
        return (Companylist) super.setOauthToken(oauthToken);
      }

      @Override
      public Companylist setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Companylist) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Companylist setQuotaUser(java.lang.String quotaUser) {
        return (Companylist) super.setQuotaUser(quotaUser);
      }

      @Override
      public Companylist setUserIp(java.lang.String userIp) {
        return (Companylist) super.setUserIp(userIp);
      }

      @com.google.api.client.util.Key
      private java.lang.String username;

      /**

       */
      public java.lang.String getUsername() {
        return username;
      }

      public Companylist setUsername(java.lang.String username) {
        this.username = username;
        return this;
      }

      @Override
      public Companylist set(String parameterName, Object value) {
        return (Companylist) super.set(parameterName, value);
      }
    }
    /**
     * Get user.
     *
     * Create a request for the method "userInfos.info".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Info#execute()} method to invoke the remote operation.
     *
     * @return the request
     */
    public Info info() throws java.io.IOException {
      Info result = new Info();
      initialize(result);
      return result;
    }

    public class Info extends InforceRequest<localhost.inforce.model.GetUserInfoResp> {

      private static final String REST_PATH = "userResult";

      /**
       * Get user.
       *
       * Create a request for the method "userInfos.info".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Info#execute()} method to invoke the remote operation. <p> {@link
       * Info#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)} must be
       * called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @since 1.13
       */
      protected Info() {
        super(Inforce.this, "GET", REST_PATH, null, localhost.inforce.model.GetUserInfoResp.class);
      }

      @Override
      public com.google.api.client.http.HttpResponse executeUsingHead() throws java.io.IOException {
        return super.executeUsingHead();
      }

      @Override
      public com.google.api.client.http.HttpRequest buildHttpRequestUsingHead() throws java.io.IOException {
        return super.buildHttpRequestUsingHead();
      }

      @Override
      public Info setAlt(java.lang.String alt) {
        return (Info) super.setAlt(alt);
      }

      @Override
      public Info setFields(java.lang.String fields) {
        return (Info) super.setFields(fields);
      }

      @Override
      public Info setKey(java.lang.String key) {
        return (Info) super.setKey(key);
      }

      @Override
      public Info setOauthToken(java.lang.String oauthToken) {
        return (Info) super.setOauthToken(oauthToken);
      }

      @Override
      public Info setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Info) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Info setQuotaUser(java.lang.String quotaUser) {
        return (Info) super.setQuotaUser(quotaUser);
      }

      @Override
      public Info setUserIp(java.lang.String userIp) {
        return (Info) super.setUserIp(userIp);
      }

      @com.google.api.client.util.Key
      private java.lang.String username;

      /**

       */
      public java.lang.String getUsername() {
        return username;
      }

      public Info setUsername(java.lang.String username) {
        this.username = username;
        return this;
      }

      @Override
      public Info set(String parameterName, Object value) {
        return (Info) super.set(parameterName, value);
      }
    }
    /**
     * Validates user and returns the companies belonged to.
     *
     * Create a request for the method "userInfos.validate".
     *
     * This request holds the parameters needed by the inforce server.  After setting any optional
     * parameters, call the {@link Validate#execute()} method to invoke the remote operation.
     *
     * @param content the {@link localhost.inforce.model.LoginReq}
     * @return the request
     */
    public Validate validate(localhost.inforce.model.LoginReq content) throws java.io.IOException {
      Validate result = new Validate(content);
      initialize(result);
      return result;
    }

    public class Validate extends InforceRequest<localhost.inforce.model.LoginResp> {

      private static final String REST_PATH = "valUserInfoResult";

      /**
       * Validates user and returns the companies belonged to.
       *
       * Create a request for the method "userInfos.validate".
       *
       * This request holds the parameters needed by the the inforce server.  After setting any optional
       * parameters, call the {@link Validate#execute()} method to invoke the remote operation. <p>
       * {@link
       * Validate#initialize(com.google.api.client.googleapis.services.AbstractGoogleClientRequest)}
       * must be called to initialize this instance immediately after invoking the constructor. </p>
       *
       * @param content the {@link localhost.inforce.model.LoginReq}
       * @since 1.13
       */
      protected Validate(localhost.inforce.model.LoginReq content) {
        super(Inforce.this, "POST", REST_PATH, content, localhost.inforce.model.LoginResp.class);
      }

      @Override
      public Validate setAlt(java.lang.String alt) {
        return (Validate) super.setAlt(alt);
      }

      @Override
      public Validate setFields(java.lang.String fields) {
        return (Validate) super.setFields(fields);
      }

      @Override
      public Validate setKey(java.lang.String key) {
        return (Validate) super.setKey(key);
      }

      @Override
      public Validate setOauthToken(java.lang.String oauthToken) {
        return (Validate) super.setOauthToken(oauthToken);
      }

      @Override
      public Validate setPrettyPrint(java.lang.Boolean prettyPrint) {
        return (Validate) super.setPrettyPrint(prettyPrint);
      }

      @Override
      public Validate setQuotaUser(java.lang.String quotaUser) {
        return (Validate) super.setQuotaUser(quotaUser);
      }

      @Override
      public Validate setUserIp(java.lang.String userIp) {
        return (Validate) super.setUserIp(userIp);
      }

      @Override
      public Validate set(String parameterName, Object value) {
        return (Validate) super.set(parameterName, value);
      }
    }

  }

  /**
   * Builder for {@link Inforce}.
   *
   * <p>
   * Implementation is not thread-safe.
   * </p>
   *
   * @since 1.3.0
   */
  public static final class Builder extends com.google.api.client.googleapis.services.json.AbstractGoogleJsonClient.Builder {

    /**
     * Returns an instance of a new builder.
     *
     * @param transport HTTP transport, which should normally be:
     *        <ul>
     *        <li>Google App Engine:
     *        {@code com.google.api.client.extensions.appengine.http.UrlFetchTransport}</li>
     *        <li>Android: {@code newCompatibleTransport} from
     *        {@code com.google.api.client.extensions.android.http.AndroidHttp}</li>
     *        <li>Java: {@link com.google.api.client.googleapis.javanet.GoogleNetHttpTransport#newTrustedTransport()}
     *        </li>
     *        </ul>
     * @param jsonFactory JSON factory, which may be:
     *        <ul>
     *        <li>Jackson: {@code com.google.api.client.json.jackson2.JacksonFactory}</li>
     *        <li>Google GSON: {@code com.google.api.client.json.gson.GsonFactory}</li>
     *        <li>Android Honeycomb or higher:
     *        {@code com.google.api.client.extensions.android.json.AndroidJsonFactory}</li>
     *        </ul>
     * @param httpRequestInitializer HTTP request initializer or {@code null} for none
     * @since 1.7
     */
    public Builder(com.google.api.client.http.HttpTransport transport, com.google.api.client.json.JsonFactory jsonFactory,
        com.google.api.client.http.HttpRequestInitializer httpRequestInitializer) {
      super(
          transport,
          jsonFactory,
          DEFAULT_ROOT_URL,
          DEFAULT_SERVICE_PATH,
          httpRequestInitializer,
          false);
    }

    /** Builds a new instance of {@link Inforce}. */
    @Override
    public Inforce build() {
      return new Inforce(this);
    }

    @Override
    public Builder setRootUrl(String rootUrl) {
      return (Builder) super.setRootUrl(rootUrl);
    }

    @Override
    public Builder setServicePath(String servicePath) {
      return (Builder) super.setServicePath(servicePath);
    }

    @Override
    public Builder setHttpRequestInitializer(com.google.api.client.http.HttpRequestInitializer httpRequestInitializer) {
      return (Builder) super.setHttpRequestInitializer(httpRequestInitializer);
    }

    @Override
    public Builder setApplicationName(String applicationName) {
      return (Builder) super.setApplicationName(applicationName);
    }

    @Override
    public Builder setSuppressPatternChecks(boolean suppressPatternChecks) {
      return (Builder) super.setSuppressPatternChecks(suppressPatternChecks);
    }

    @Override
    public Builder setSuppressRequiredParameterChecks(boolean suppressRequiredParameterChecks) {
      return (Builder) super.setSuppressRequiredParameterChecks(suppressRequiredParameterChecks);
    }

    @Override
    public Builder setSuppressAllChecks(boolean suppressAllChecks) {
      return (Builder) super.setSuppressAllChecks(suppressAllChecks);
    }

    /**
     * Set the {@link InforceRequestInitializer}.
     *
     * @since 1.12
     */
    public Builder setInforceRequestInitializer(
        InforceRequestInitializer inforceRequestInitializer) {
      return (Builder) super.setGoogleClientRequestInitializer(inforceRequestInitializer);
    }

    @Override
    public Builder setGoogleClientRequestInitializer(
        com.google.api.client.googleapis.services.GoogleClientRequestInitializer googleClientRequestInitializer) {
      return (Builder) super.setGoogleClientRequestInitializer(googleClientRequestInitializer);
    }
  }
}
