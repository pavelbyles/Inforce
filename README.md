Just testing ...

1. Appengine
2. Testing (TDD)
3. Templates UI
4. Bootstrap
5. Less

`sudo npm install -g bower`<br />
`sudo npm install -g vulcanize`<br />
`bower init`<br />
`bower install jquery`<br />
`bower install font-awesome`<br />
`bower install bootswatch`<br />
`bower install less<`br />
`bower install Happyjs`<br />
`bower install polymer Polymer/polymer-elements Polymer/polymer-ui-elements`<br />
`bower install Polymer/core-ajax<`br />
<br />
View / test api: `http://localhost:8080/_ah/api/explorer`<br />

#Generate Android client libs:
URL='http://localhost:8080/_ah/api/discovery/v1/apis/company/v1/rest'<br />
`> curl -s $URL > inforce.rest.discovery`
`> $GO_SDK/endpointscfg.py gen_client_lib java inforce.rest.discovery`
`> unzip inforce.rest.zip`

// npm update -g bower
